type parser = JSON of (int list) * string | PARSER | AST | JSONIFY of { mutable fields : int; other : string }

let get_parser (p : parser) =
  match p with 
  | JSON (l, s)                                 -> s
  | PARSER                                      -> "PARSER"
  | AST                                         -> "AST"
  | JSONIFY { fields = temp; other = value; }   -> value
;;

let update_parser (x : parser) = 
  match x with 
  | JSONIFY a -> a.fields <- a.fields + 200
  | _ -> ()
;;

let mutability (value : int ref) = 
  (value := 2000); 
  Printf.printf "%i" !value;
;;   

let rec get_tail_and_head (a : int list) = 
  match a with 
  | head :: tail -> if tail == [] then (head :: []) else get_tail_and_head tail
  | _ -> []
;;

let concat (a : int list) = [ 2 ] @ a @ (get_tail_and_head a)
;;


let main = 
  let 
    p = JSONIFY { fields = 12; other = "aman" };
    in  
      let 
        _ = update_parser p;
          in
            match p with
            | JSONIFY a -> Printf.printf "%i" a.fields
            | _         -> ()
;;

main

