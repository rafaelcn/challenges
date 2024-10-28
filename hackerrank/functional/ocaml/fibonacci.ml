let modulo = 10000007L;;

(*
  Is the modulo of fibonacci of a nth number equals to the fibonacci of the nth
  number? I need to figure out an identity that gives me the same result to work
  with Int64.

  setting aside the differences when computing modulo in OCaml lets say that the
  % operator is the modulo operator and it doesn't have any type width constraint.
  
  1e8+7 % Fib(10) = 55
  1e8+7 % Fib(100) = 24278230

  1e8+7 % Fib(100) = ???
  
*)

let fibonacci n =
  if n = 0L then
    0L
  else
    let rec tail m (acc1: int64) (acc2: int64) =
      Printf.printf "n: %Ld acc1: %Ld acc2: %Ld\n" m acc1 acc2;
      if m = 2L then
        Int64.add acc1 acc2
      else
        tail (Int64.sub m 1L) acc2 (Int64.add acc1 acc2)
    in tail n 0L 1L;;

let solve n =
  let r = Int64.rem (fibonacci n) modulo in
  Printf.printf "%Ld\n" r;;

let () =
  let _ = read_int() in (* ignore first input *)
  let rec reader l =
      try
        let i = (Int64.of_int (read_int ())) in
        reader (i :: l)
      with
      | End_of_file -> List.rev l
  in
  let numbers = reader [] in
  List.iter solve numbers;;
