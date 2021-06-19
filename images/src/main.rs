use std::io::Write;

use actix_files::NamedFile;
use actix_multipart::Multipart;
use actix_web::{self, get, post, web, App, Error, HttpResponse, HttpServer, Result};
use futures::{StreamExt, TryStreamExt};
use serde::Deserialize;
use std::path::PathBuf;

#[derive(Deserialize)]
struct FileParams {
    id: String,
    filename: String,
}

#[post("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}")]
async fn save_file(
    mut payload: Multipart,
    data: web::Path<FileParams>,
) -> Result<HttpResponse, Error> {
    while let Ok(Some(mut field)) = payload.try_next().await {
        let filepath = format!(
            "./imagestore/{}",
            data.id,
        );

        let mut path = PathBuf::from(filepath);
        if !path.exists() {
            std::fs::create_dir(&path)?;
            path = path.join(PathBuf::from(&data.filename));
            let mut f = web::block(|| std::fs::File::create(path))
                .await
                .unwrap();
    
            while let Some(chunk) = field.next().await {
                let data = chunk.unwrap();
                f = web::block(move || f.write_all(&data).map(|_| f)).await?;
            }
        } else {
            path = path.join(PathBuf::from(&data.filename));
            std::fs::remove_file(path.file_name().unwrap())?;
            let mut f = web::block(|| std::fs::File::create(path))
            .await
            .unwrap();

            while let Some(chunk) = field.next().await {
                let data = chunk.unwrap();
                f = web::block(move || f.write_all(&data).map(|_| f)).await?;
            }
        }
    }

    Ok(HttpResponse::Ok().into())
}

#[get("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}")]
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
        .service(get_file)
        .service(save_file))
        .bind("127.0.0.1:8080")?
        .run()
        .await
}
