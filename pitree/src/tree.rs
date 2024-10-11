struct TreeNode2 {
    nodeid: u8,
    level: u16,
    children: Vec<TreeNode2>
} 

impl TreeNode2 {
    fn new(id: u8, level: u16) -> Self {
        TreeNode2 { nodeid: id, level, children: Vec::new() }
    } 

    fn build(&mut self) -> () {
        
    }
}