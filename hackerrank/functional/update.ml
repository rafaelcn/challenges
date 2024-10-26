let update l =
  let rec curry l u =
  match l with
  | [] -> u
  | h :: t -> curry t (abs h :: u) in
  curry l [];;

let () =
  let rec reader l = 
    try
      let n = read_int () in
      reader (n :: l)
    with End_of_file -> l
  in
  let numbers = reader [] in
  List.iter (Printf.printf "%d\n") (update numbers);;
