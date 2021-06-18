use std::io;

pub trait Storage {
    fn save(&self, path: String, file: &mut Box<dyn io::Read>) -> io::Result<()>;
}