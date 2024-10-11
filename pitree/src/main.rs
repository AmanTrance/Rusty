mod tree;
use std::f64::consts::PI;

#[derive(Debug, PartialEq, PartialOrd)]
struct TreeNode {
    nodeid: u8,
    level: u16,
    children: Vec<TreeNode>,
    ptr: Option<Box<TreeNode>>
}

impl TreeNode {
    fn new(id: u8, level: u16) -> Self {
        TreeNode { nodeid: id, level, children: Vec::new(), ptr: None }
    }

    fn build(&mut self) -> () {
        let mut nodeid: u8 = self.nodeid + 1;
        let mut level: u16 = self.level + 1;
        let mut ptr: &mut TreeNode = self;
        for i in pi_values().into_iter() {
            for _ in 0..i.to_digit(10).unwrap() {
                let node: TreeNode = Self::new(nodeid, level);
                nodeid += 1;
                ptr.children.push(node);
            }
            level += 1;
            ptr.ptr = Some(Box::new(Self::new(nodeid, level)));
            nodeid += 1;
            ptr = ptr.ptr.as_mut().unwrap();
        }
    }

    fn traverse(&self) -> () {
        let mut root: &TreeNode = self;
        while root.ptr != None {
            print!("{} ", root.nodeid);
            for child in root.children.iter() {
                print!("{} ", child.nodeid);
            }
            root = root.ptr.as_ref().unwrap();    
        }
        print!("{}", root.nodeid); 
    }
}

fn pi_values() -> Vec<char> {
    let pi: f64 = PI;
    let pi_string: String = format!("{}", pi);
    let mut values: Vec<char> = Vec::new();
    for (index, i) in pi_string.chars().into_iter().enumerate() {
        if index == 0 || index == 1 || index == 2 { continue; }
        values.push(i);
    }
    values
}

fn main() {
    let mut tree: TreeNode = TreeNode::new(1, 1);
    tree.build();
    tree.traverse();
    println!("\n{tree:?}");
}


