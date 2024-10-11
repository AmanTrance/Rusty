use std::fmt::Display;

#[allow(dead_code)]
#[derive(Debug,Clone,PartialEq, PartialOrd)]
pub struct Queue{
    pub data: Vec<i32>
}
#[allow(dead_code)]
impl Queue {
    pub fn new(x: Vec<i32>) -> Self {
        Self { data: x }
    }
    pub fn append(&mut self, x: i32) -> () {
        self.data.push(x);
    }
    pub fn popfront(&mut self) -> () {
        let mut z: Vec<i32> = Vec::new();
        for i in self.data.iter(){
            if *i == self.data[0]{
                continue;
            }
            else{
                z.push(*i);
            }
        }
        self.data = z;
    }
}
impl Display for Queue {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{:?}", self.data)
    }
}
