% Base de conocimiento
carrera(ingenieria, sistemas, logica, programacion, tecnologia).
carrera(ingenieria, civil, matematica, dibujo, construccion).
carrera(medicina, medicina_general, biologia, empatia, salud).
carrera(humanidades, filosofia, analisis, redaccion, lectura).

% Inferencia
recomendar_carrera(Apt, Hab, Int, Fac, Car) :-
    carrera(Fac, Car, Apt, Hab, Int).

% Instrucción automática
:- initialization(main).

main :-
    recomendar_carrera(logica, programacion, tecnologia, F, C),
    format("Facultad: ~w~n", [F]),
    format("Carrera: ~w~n", [C]).
