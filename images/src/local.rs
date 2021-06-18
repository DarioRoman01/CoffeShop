use std::io::{Error, ErrorKind};
use std::io;
use std::fs::{self, File};
use std::path::PathBuf;
use crate::storage::Storage;

pub struct Local {
    pub max_size: i32,
    pub base_path: PathBuf,
}

impl Local {
    pub fn new(path: String, size: i32) -> io::Result<Local> {
        let buf = PathBuf::from(path);
        let p = fs::canonicalize(&buf);

        match p {
            Ok(e) => return Ok(Local {max_size: size, base_path: e}),
            Err(_) => todo!(),
        }
    }

    fn full_path(&self, path: String) -> PathBuf {
        let p = PathBuf::from(path);
        return self.base_path.join(p);
    }

    pub fn get(&self, path: String) -> io::Result<fs::File> {
        let fp = self.full_path(path);

        let file = match File::open(&fp) {
            Ok(file) => file,
            Err(e) => return Err(e),
        };

        return Ok(file);
    }
}

impl Storage for Local {
    fn save(&self, path: String, file: &mut Box<dyn io::Read>) -> io::Result<()> {
        let fp = self.full_path(path);
        let dir = fp.parent().unwrap();

        match fs::create_dir(dir) {
            Ok(()) => (),
            Err(e) => return Err(e),
        };

        if fp.exists() {
            match fs::remove_file(fp.clone()) {
                Ok(()) => (),
                Err(e) => return Err(e),
            };
        } else {
            return Err(Error::new(ErrorKind::NotFound, "Unable to get file info"));
        }

        let mut f = fs::File::create(fp)?;
        io::copy(file, &mut f)?;
        return Ok(())
    }
}
