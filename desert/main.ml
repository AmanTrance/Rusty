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

type explicit = E | H 
  and t = L | P of explicit * int
;;
 
class haha (x : int)=
  let y = x + 100 in
  object(self)
    val mutable hello : int = y
    method change : unit = Printf.printf "%s" "aman"
end
;;

class derived s= 
  object(self)
  inherit haha s
  val mutable x : int = 0
  method copy = {< x = x + 200 >} 
end 
;;

let (--=@) : (int -> int) -> int = function 
  | a -> a 2000
;;

let () =
  let d = new derived 100 in 
    d#copy#change;
    Printf.printf "%i" @@ (--=@) @@ (fun x: int -> x);
;;

let k = 
  let v = 90 + 100 in 
  v
;;

let y (z : int option) = function 
  | z -> match z with 
        | Some a  -> if a == 100 then a else 0 
        | None    -> 0x1f
;;

let get_array = let 
  a = ref [| 1; 2; 3; |]
    in
      a := Array.append !a [| 3; 4; |] ;
      !a
;;
 
class ['a] h =
  object(self)
  val first : 'a option = None
  method bind_first : 'a = match first with 
                            | Some x  -> 100
                            | None    -> 200  
end
;;

let jj (a : int) = 
  let y = a + 10 and x = a + 10
    in
    y + x
;;

let (let*) (i : int) (f : int -> int) = 
  match i with 
    | 10  -> f i 
    | _   -> f i
;;

Printf.printf "%i" @@ jj 20


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

