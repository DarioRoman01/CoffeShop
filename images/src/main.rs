use std::io::Write;

use actix_files::NamedFile;
use actix_cors::Cors;
use actix_multipart::Multipart;
use actix_web::{self, web, App, Error, HttpResponse, HttpServer, Result, http::header, middleware};
use futures::{StreamExt, TryStreamExt};
use serde::Deserialize;
use std::{fs, path::PathBuf};

#[derive(Deserialize)]
struct FileParams {
    id: String,
    filename: String,
}

async fn save_file(
    mut payload: Multipart,
    web::Path(id): web::Path<u32>
) -> Result<HttpResponse, Error> {
    while let Ok(Some(mut field)) = payload.try_next().await {
        let content_type = field.content_disposition().unwrap();
        let filename = content_type.get_filename().unwrap();
        let filepath = format!("./imagestore/{}", &id);

        let mut path = PathBuf::from(filepath);
        if !path.exists() {
            fs::create_dir(&path)?;
            path = path.join(PathBuf::from(sanitize_filename::sanitize(&filename)));
            let mut f = web::block(|| fs::File::create(path))
                .await
                .unwrap();
    
            while let Some(chunk) = field.next().await {
                let data = chunk.unwrap();
                f = web::block(move || f.write_all(&data).map(|_| f)).await?;
            }
        } else {
            path = path.join(PathBuf::from(sanitize_filename::sanitize(&filename)));
            match fs::remove_file(&path) {
                Ok(()) => (),
                Err(_) => return Ok(HttpResponse::NotFound().body("Unable to update the requested file")),
            }

            let mut f = web::block(|| fs::File::create(path))
            .await
            .unwrap();

            while let Some(chunk) = field.next().await {
                let data = chunk.unwrap();
                f = web::block(move || f.write_all(&data).map(|_| f)).await?;
            }
        }
    }

    Ok(HttpResponse::Ok().body("File created succesfully"))
}

async fn delete_file(file_data: web::Path<FileParams>) -> HttpResponse {
    let path = PathBuf::from(format!(
        "./imagestore/{}/{}",
        file_data.id, file_data.filename
    ));

    match fs::remove_file(path) {
        Ok(_) => return HttpResponse::Ok().body("file deleted succesfully"),
        Err(_) => return HttpResponse::NotFound().body("File does not exist"),
    }
}

async fn get_file(file_data: web::Path<FileParams>) -> Result<NamedFile> {
    let path = PathBuf::from(format!(
        "./imagestore/{}/{}",
        file_data.id, file_data.filename
    ));

    Ok(NamedFile::open(path)?)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| 
        App::new()
        .wrap(middleware::Compress::default())
        .wrap(Cors::default()
            .allowed_origin("http://localhost:3000")
            .allowed_methods(vec!["GET", "POST", "DELETE"])
            .allowed_header(header::CONTENT_TYPE)
            .max_age(3600)
        )
        .service(
            web::resource("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}")
            .route(web::get().to(get_file))
            .route(web::delete().to(delete_file))
        )
        .service(
            web::resource("/images/{id:[0-9]+}")
            .route(web::post().to(save_file))
        ))
        .bind("127.0.0.1:8080")?
        .run()
        .await
}