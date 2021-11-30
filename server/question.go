package server

type Question struct {
	questionNumber int
	questionStr    string
	options        []string
	correctOption  int
}

func CreateRandomQuestions() []Question {
	questions := []Question{
		{
			questionStr:   "Al aprobar la cursada de una materia cuantos cuatrimestres tiene uno para poder rendir el final ?",
			options:       []string{"1. ", "6. ", "3. "},
			correctOption: 3},
		{
			questionStr:   "Al aprobar la cursada de una materia cuantas posibilidades tengo para rendir el final? ",
			options:       []string{"1. ", "3. ", "Depende de cada profesor. "},
			correctOption: 2},
		{
			questionStr:   "Segun la reciente actualización del plan de estudios de Ingeniería Informatica , Cuál de las siguientes materias no fue sacada como obligatoria para la carrera? ",
			options:       []string{"Quimica. ", "Analisis Matematico III. ", "Fisica II. "},
			correctOption: 3},
		{
			questionStr:   "Segun las encuestas del departamento de computación , cual de las siguientes cátedras de la materia Algoritmos y Programación II es la mas elegida por los estudiantes al momento de inscribirse? ",
			options:       []string{"Calvo. ", "Mendez. ", "Juarez. "},
			correctOption: 2},
		{
			questionStr:   "Segun el calendario académico , cuantas semanas hay para rendir finales ? ",
			options:       []string{"2. ", "4. ", "7. "},
			correctOption: 3},
		{
			questionStr:   "Que es lo que se conoce como ‘Prioridad’?  ",
			options:       []string{"Una materia. ", "Un numero que indica el orden en que los estudiantes pueden inscribirse favoreciendo así a los que mejor prioridad tienen a conseguir cupo en las mejores cátedras. ", "Ninguna de las anteriores es correcta. "},
			correctOption: 2},
		{
			questionStr:   "Quien tiene mejor prioridad ? ",
			options:       []string{"Aquel alumno con un numero lo mas grande posible. ", "Aquel alumno con un numero lo mas pequeño  posible. ", "Depende. "},
			correctOption: 2},
		{
			questionStr:   "Que significa estar como alumno ‘condicional’ en una materia ?  ",
			options:       []string{"Que no vas a poder cursar la materia. ", "Que como todas las cátedras habían llenado sus cupos , el departamento de computación te asigna una cátedra aleatoria. ", "Que todavía no aprobaste el final de la materia. "},
			correctOption: 2},
		{
			questionStr:   "En el ámbito de la informática cual fue el plan que se actualizo recientemente ? ",
			options:       []string{"El Plan de Ingeniería informática. ", "El plan de la licenciatura en sistemas. ", "Ambos. "},
			correctOption: 1},
		{
			questionStr:   "En donde se encuentra la FIUBA ? ",
			options:       []string{"Av. Paseo Colón 850. ", "Av. Gral. Las Heras 2214. ", "Ambas opciones son correctas. "},
			correctOption: 3},
		{
			questionStr:   "Cual es el horario en el que suelen cursarse las materias informáticas?  ",
			options:       []string{"A la mañana. ", "Al mediodia. ", "A la noche. "},
			correctOption: 3},
		{
			questionStr:   "Cual es la catedra preferida por lo alumnos al momento de cursar física I ? ",
			options:       []string{"Fontana. ", "Cornejo. ", "Garea. "},
			correctOption: 1},
		{
			questionStr:   "Segun el plan de ingeniería informática cual es la duración de dicha carrera ? ",
			options:       []string{"2. ", "4. ", "6. "},
			correctOption: 3},
		{
			questionStr:   "De acuerdo con las estadísticas de la facultad, en cuanto tiempo logra recibirse el estudiante promedio de ingeniería informática??",
			options:       []string{"4. ", "6. ", "9. "},
			correctOption: 3},
		{
			questionStr:   " Que es el SIU guarani  ?",
			options:       []string{"El sistema de autogestion de la facultad. ", "Un aula de la facultad. ", "Una lengua indigena. "},
			correctOption: 1},
		{
			questionStr:   "Generalmente cuando son las inscripciones a las materias ?",
			options:       []string{"2 Semanas antes del comienzo de clases.  ", "1 Semana antes del comienzo de clases. ", "5 semanas antes del comienzo de clases. "},
			correctOption: 2},
		{
			questionStr:   "Si apruebo el final de una materia donde puedo encontrar la nota en el sistema ? ",
			options:       []string{"En inscripciones a materias. ", "En inscripciones a exámenes. ", "En actuación provisoria. "},
			correctOption: 3},
		{
			questionStr:   "Cual es la catedra preferida por los estudiantes al momento de cursar algebra II?",
			options:       []string{"Grynberg. ", "Lopez. ", "Alvarez. "},
			correctOption: 1},
		{
			questionStr:   "Segun el plan de estudios de Ing Informatica, cual de las siguientes materias no es correlativa de la materia Teoria de lenguaje ?  ",
			options:       []string{"Algoritmos y Programación II. ", "Algoritmos y Programación I. ", "Fisica II. "},
			correctOption: 1},
		{
			questionStr:   "Que pasa si me anoto en una materia y luego decido no cursarla? ",
			options:       []string{"Te afecta el promedio. ", "Por 2 cuartrimestres no te podes volver a anotar. ", "Nada. "},
			correctOption: 3},
	}
	return questions
}
