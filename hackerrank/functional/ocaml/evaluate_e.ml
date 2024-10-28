let power base exp =
  let rec curry base exp acc =
    if exp > 1 then
      curry base (exp-1) acc*.base
    else
      base
  in curry base exp 1.;;

let factorial n =
  let rec curry n acc =
    if n == 0 then
      acc
    else
      curry (n-1) (n*acc)
  in
  curry n 1;;


let evaluate n =
  let rec curry x n res =
    if x >= 1 then
      let f = (float_of_int (factorial x)) in
      let p = (power n x) in
      (*Printf.printf "%d -> %f %f\n" x f p;*)
      curry (x-1) n (res +. p/.f)
  else
    res
  in
  curry 9 n 1.;;

let () =
  let _ = read_float () in (* ignore the first input *)
  let rec reader l = 
    try
      let n = read_float () in
      reader (n :: l)
    with End_of_file -> List.rev l
  in
  let numbers = reader [] in
  let evaluations = List.map evaluate numbers in
  List.iter (Printf.printf "%.4f\n" ) evaluations;;

