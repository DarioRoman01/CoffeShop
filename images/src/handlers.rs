use crate::storage::Storage;

pub struct Files {
    store: Box<dyn Storage>,
}

impl Files {
    pub fn new(s: Box<dyn Storage>) -> Self {
        return Files{store: s};
    }
}