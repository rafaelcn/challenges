let rec sum_odd l acc =
  match l with
  | [] -> acc
  | h :: t -> if Int.abs (h mod 2) == 1 then sum_odd t (acc+h) else sum_odd t acc;;

let () =
  let rec reader l =
    try
      let n = read_int () in
      reader (n :: l)
    with End_of_file -> l
  in
  let numbers = reader [] in
  Printf.printf "%d\n" (sum_odd numbers 0);;
