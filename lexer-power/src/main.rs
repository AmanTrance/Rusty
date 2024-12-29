use std::{cell::{RefCell, RefMut}, rc::Rc};

fn closure_return_without_gc<'a, 'b, T, F>(x: &'a Rc<RefCell<T>>, operation: fn(RefMut<'_, T>, F) -> ()) 
    -> impl Fn(F) -> Rc<RefCell<T>> + 'b
        where 
            'a      : 'b,
            'static : 'a,
            'static : 'b,
            F       : 'a
    {

    return move |a: F| {
        let captured_value: Rc<RefCell<T>> = Rc::clone(x);
        operation(captured_value.borrow_mut(), a);
        captured_value
    }
}

fn generator() -> impl Fn() -> i32  {
    let value: Rc<RefCell<i32>> = Rc::new(RefCell::new(10));

    return move || {
        let captured_value: Rc<RefCell<i32>> = Rc::clone(&value);
        *captured_value.borrow_mut() += 1;
        let return_value: i32 = *captured_value.borrow();
        return_value
    }
}

fn main() {
    let k: Rc<RefCell<String>> = Rc::new(RefCell::new(String::new()));
    let value = closure_return_without_gc(&k, |mut a: RefMut<'_, String>, b: &str| { a.push_str(b); });
    println!("{}", value("haha").borrow());
    println!("{}", value("again").borrow());

    let g = generator();
    println!("{}", g());
    println!("{}", g());
    println!("{}", g());
}
