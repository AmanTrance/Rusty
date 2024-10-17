use std::hash::{DefaultHasher, Hash, Hasher};

pub struct BloomFilter {
    filter: u32,
    pub size: u32
}

impl BloomFilter {
    pub fn new(size: u32) -> Self {
        BloomFilter { filter: 0, size }
    }

    pub fn hash_and_store(&mut self, value: String) -> () {
        let mut hasher = DefaultHasher::new();
        value.hash(&mut hasher);
        let temp = hasher.finish() as u32;
        self.filter |= temp;
    }

    pub fn contains(&self, value: String) -> bool {
        let mut hasher = DefaultHasher::new();
        value.hash(&mut hasher);
        let temp = hasher.finish() as u32;
        if (self.filter & temp) == temp {
            true
        } else {
            false
        }
    }
}

