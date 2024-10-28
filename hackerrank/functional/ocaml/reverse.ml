(*
  This solution is a bit weird as the reverse function is not used. This is the
  effect of reading the input that naturally reverses the input order, I could use
  List.rev in the result of the input reading but I chose not to as this would add
  things to the program that doesn't need to be added.
*)

let reverse l =
  let rec curry l acc =
    match l with
    | [] -> acc
    | h :: t -> curry t (h :: acc) in
  curry l [];;

let () =
  let rec reader l = 
    try
      let n = read_int () in
      reader (n :: l)
    with End_of_file -> l
  in
  let numbers = reader [] in
  List.iter (Printf.printf "%d\n") numbers;;
