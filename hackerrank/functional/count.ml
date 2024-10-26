let count l =
  let rec curry l acc =
    match l with
    | [] -> acc
    | h :: t -> curry t (acc+1)
  in
  curry l 0;;

let () =
  let rec reader l = 
    try
      let n = read_int () in
      reader (n :: l)
    with End_of_file -> l
  in
  let numbers = reader [] in
 Printf.printf "%d\n" (count numbers);;
