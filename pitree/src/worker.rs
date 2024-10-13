use std::collections::{BinaryHeap, HashMap};
use float_ord::FloatOrd;
use core::f64;

#[allow(dead_code)]
struct Solution{}
impl Solution {
    fn mincost_to_hire_workers(self, quality: Vec<i32>, wage: Vec<i32>, k: i32) -> f64 {
        let mut ratio_vec: Vec<f64> = Vec::new();
        for i in 0..quality.len(){
            let temp: f64 = wage[i] as f64 / quality[i] as f64;
            ratio_vec.push(temp);
        }
        let mut map: HashMap<FloatOrd<f64>, i32> = HashMap::new();
        for i in 0..ratio_vec.len(){
            map.insert(FloatOrd(ratio_vec[i]), quality[i]);
        }      
        float_ord::sort(&mut ratio_vec);
        let mut heap: BinaryHeap<i32> = BinaryHeap::new();
        let mut curr_pay: f64 = f64::MAX;
        for i in ratio_vec.into_iter(){
            if heap.len() != k as usize{
                heap.push(map.get(&FloatOrd(i)).unwrap().to_owned());
            }
            if heap.len() == k as usize{
                let temp_vec: Vec<i32> = heap.clone().into_vec();
                let mut tempf: f64 = 0.0;
                for j in temp_vec.into_iter(){
                    tempf += j as f64 * i;
                }
                if tempf < curr_pay{
                    curr_pay = tempf;
                }
                heap.pop();
            }
        }
        curr_pay

        
    }
}
fn main(){
    let quality = vec![3,1,10,10,1];
    let wage = vec![4,8,2,2,7];
    let k = 3;
    let temp = Solution{};
    let ans = temp.mincost_to_hire_workers(quality, wage, k);
    println!("{ans}");
}