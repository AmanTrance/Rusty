use crate::pi_values;
pub struct TreeNode2 {
    nodeid: u8,
    level: u16,
    children: Vec<TreeNode2>
} 

impl TreeNode2 {
    pub fn new(id: u8, level: u16) -> Self {
        TreeNode2 { nodeid: id, level, children: Vec::new() }
    } 

    pub fn build(&mut self) -> () {
        let mut nodeid: u8 = self.nodeid + 1;
        let mut level: u16 = self.level + 1;
        let mut ptr: &mut TreeNode2 = self;
        for i in pi_values().into_iter() {
            for _ in 0..i.to_digit(10).unwrap() {
                let node: TreeNode2 = TreeNode2 { nodeid, level, children: Vec::new() };
                ptr.children.push(node);
                nodeid += 1;
            }
            level += 1;
            ptr = &mut ptr.children[0];
        }
    }

    pub fn traverse(&self) {
        let mut ptr: &TreeNode2 = self;
        print!("{} ", ptr.nodeid);
        while ptr.children.len() != 0 {
            for i in ptr.children.iter() {
                print!("{} ", i.nodeid);
            }
            ptr = &ptr.children[0];
        }
    }
}