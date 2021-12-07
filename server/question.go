package server

type Question struct {
	questionNumber int
	questionStr    string
	options        []string
	correctOption  int
}

//returns a random question for players to answer
func CreateRandomQuestions() []Question {
	questions := []Question{
		{
			questionStr:   "A̳l̳ a̳p̳r̳o̳b̳a̳r̳ l̳a̳ c̳u̳r̳s̳a̳d̳a̳ d̳e̳ u̳n̳a̳ m̳a̳t̳e̳r̳i̳a̳ c̳u̳a̳n̳t̳o̳s̳ c̳u̳a̳t̳r̳i̳m̳e̳s̳t̳r̳e̳s̳ t̳i̳e̳n̳e̳ u̳n̳o̳ p̳a̳r̳a̳ p̳o̳d̳e̳r̳ r̳e̳n̳d̳i̳r̳ e̳l̳ f̳i̳n̳a̳l̳ ?̳",
			options:       []string{"1. ", "6. ", "3. "},
			correctOption: 3},
		{
			questionStr:   "A̳l̳ a̳p̳r̳o̳b̳a̳r̳ l̳a̳ c̳u̳r̳s̳a̳d̳a̳ d̳e̳ u̳n̳a̳ m̳a̳t̳e̳r̳i̳a̳ c̳u̳a̳n̳t̳a̳s̳ p̳o̳s̳i̳b̳i̳l̳i̳d̳a̳d̳e̳s̳ t̳e̳n̳g̳o̳ p̳a̳r̳a̳ r̳e̳n̳d̳i̳r̳ e̳l̳ f̳i̳n̳a̳l̳?̳ ",
			options:       []string{"1. ", "3. ", "Depende de cada profesor. "},
			correctOption: 2},
		{
			questionStr:   "S̳e̳g̳u̳n̳ l̳a̳ r̳e̳c̳i̳e̳n̳t̳e̳ a̳c̳t̳u̳a̳l̳i̳z̳a̳c̳i̳ó̳n̳ d̳e̳l̳ p̳l̳a̳n̳ d̳e̳ e̳s̳t̳u̳d̳i̳o̳s̳ d̳e̳ I̳n̳g̳e̳n̳i̳e̳r̳í̳a̳ I̳n̳f̳o̳r̳m̳a̳t̳i̳c̳a̳ ,̳ C̳u̳á̳l̳ d̳e̳ l̳a̳s̳ s̳i̳g̳u̳i̳e̳n̳t̳e̳s̳ m̳a̳t̳e̳r̳i̳a̳s̳ n̳o̳ f̳u̳e̳ s̳a̳c̳a̳d̳a̳ c̳o̳m̳o̳ o̳b̳l̳i̳g̳a̳t̳o̳r̳i̳a̳ p̳a̳r̳a̳ l̳a̳ c̳a̳r̳r̳e̳r̳a̳?̳ ",
			options:       []string{"Quimica. ", "Analisis Matematico III. ", "Fisica II. "},
			correctOption: 3},
		{
			questionStr:   "S̳e̳g̳u̳n̳ l̳a̳s̳ e̳n̳c̳u̳e̳s̳t̳a̳s̳ d̳e̳l̳ d̳e̳p̳a̳r̳t̳a̳m̳e̳n̳t̳o̳ d̳e̳ c̳o̳m̳p̳u̳t̳a̳c̳i̳ó̳n̳ ,̳c̳u̳a̳l̳ d̳e̳ l̳a̳s̳ s̳i̳g̳u̳i̳e̳n̳t̳e̳s̳ c̳á̳t̳e̳d̳r̳a̳s̳ d̳e̳ l̳a̳ m̳a̳t̳e̳r̳i̳a̳ A̳l̳g̳o̳r̳i̳t̳m̳o̳s̳ y̳ P̳r̳o̳g̳r̳a̳m̳a̳c̳i̳ó̳n̳ I̳I̳ e̳s̳ l̳a̳ m̳a̳s̳ e̳l̳e̳g̳i̳d̳a̳ p̳o̳r̳ l̳o̳s̳ e̳s̳t̳u̳d̳i̳a̳n̳t̳e̳s̳ a̳l̳ m̳o̳m̳e̳n̳t̳o̳ d̳e̳ i̳n̳s̳c̳r̳i̳b̳i̳r̳s̳e̳?̳",
			options:       []string{"Calvo. ", "Mendez. ", "Juarez. "},
			correctOption: 2},
		{
			questionStr:   "S̳e̳g̳u̳n̳ e̳l̳ c̳a̳l̳e̳n̳d̳a̳r̳i̳o̳ a̳c̳a̳d̳é̳m̳i̳c̳o̳ ,̳ c̳u̳a̳n̳t̳a̳s̳ s̳e̳m̳a̳n̳a̳s̳ h̳a̳y̳ p̳a̳r̳a̳ r̳e̳n̳d̳i̳r̳ f̳i̳n̳a̳l̳e̳s̳ ?̳ ",
			options:       []string{"2. ", "4. ", "7. "},
			correctOption: 3},
		{
			questionStr:   "Q̳u̳e̳ e̳s̳ l̳o̳ q̳u̳e̳ s̳e̳ c̳o̳n̳o̳c̳e̳ c̳o̳m̳o̳ ‘̳P̳r̳i̳o̳r̳i̳d̳a̳d̳’̳?̳  ",
			options:       []string{"Una materia. ", "Un numero que indica el orden en que los estudiantes pueden inscribirse favoreciendo así a los que mejor prioridad tienen a conseguir cupo en las mejores cátedras. ", "Ninguna de las anteriores es correcta. "},
			correctOption: 2},
		{
			questionStr:   "Q̳u̳i̳e̳n̳ t̳i̳e̳n̳e̳ m̳e̳j̳o̳r̳ p̳r̳i̳o̳r̳i̳d̳a̳d̳ ?̳ ",
			options:       []string{"Aquel alumno con un numero lo mas grande posible. ", "Aquel alumno con un numero lo mas pequeño  posible. ", "Depende. "},
			correctOption: 2},
		{
			questionStr:   "Q̳u̳e̳ s̳i̳g̳n̳i̳f̳i̳c̳a̳ e̳s̳t̳a̳r̳ c̳o̳m̳o̳ a̳l̳u̳m̳n̳o̳ ‘̳c̳o̳n̳d̳i̳c̳i̳o̳n̳a̳l̳’̳ e̳n̳ u̳n̳a̳ m̳a̳t̳e̳r̳i̳a̳ ?̳",
			options:       []string{"Que no vas a poder cursar la materia. ", "Que como todas las cátedras habían llenado sus cupos , el departamento de computación te asigna una cátedra aleatoria. ", "Que todavía no aprobaste el final de la materia. "},
			correctOption: 2},
		{
			questionStr:   "E̳n̳ e̳l̳ á̳m̳b̳i̳t̳o̳ d̳e̳ l̳a̳ i̳n̳f̳o̳r̳m̳á̳t̳i̳c̳a̳ c̳u̳a̳l̳ f̳u̳e̳ e̳l̳ p̳l̳a̳n̳ q̳u̳e̳ s̳e̳ a̳c̳t̳u̳a̳l̳i̳z̳o̳ r̳e̳c̳i̳e̳n̳t̳e̳m̳e̳n̳t̳e̳ ?̳ ",
			options:       []string{"El Plan de Ingeniería informática. ", "El plan de la licenciatura en sistemas. ", "Ambos. "},
			correctOption: 1},
		{
			questionStr:   "E̳n̳ d̳o̳n̳d̳e̳ s̳e̳ e̳n̳c̳u̳e̳n̳t̳r̳a̳ l̳a̳ F̳I̳U̳B̳A̳ ?̳ ",
			options:       []string{"Av. Paseo Colón 850. ", "Av. Gral. Las Heras 2214. ", "Ambas opciones son correctas. "},
			correctOption: 3},
		{
			questionStr:   "C̳u̳a̳l̳ e̳s̳ e̳l̳ h̳o̳r̳a̳r̳i̳o̳ e̳n̳ e̳l̳ q̳u̳e̳ s̳u̳e̳l̳e̳n̳ c̳u̳r̳s̳a̳r̳s̳e̳ l̳a̳s̳ m̳a̳t̳e̳r̳i̳a̳s̳ i̳n̳f̳o̳r̳m̳á̳t̳i̳c̳a̳s̳?̳",
			options:       []string{"A la mañana. ", "Al mediodia. ", "A la noche. "},
			correctOption: 3},
		{
			questionStr:   "C̳u̳a̳l̳ e̳s̳ l̳a̳ c̳a̳t̳e̳d̳r̳a̳ p̳r̳e̳f̳e̳r̳i̳d̳a̳ p̳o̳r̳ l̳o̳ a̳l̳u̳m̳n̳o̳s̳ a̳l̳ m̳o̳m̳e̳n̳t̳o̳ d̳e̳ c̳u̳r̳s̳a̳r̳ f̳í̳s̳i̳c̳a̳ I̳ ?̳ ",
			options:       []string{"Fontana. ", "Cornejo. ", "Garea. "},
			correctOption: 1},
		{
			questionStr:   "S̳e̳g̳u̳n̳ e̳l̳ p̳l̳a̳n̳ d̳e̳ i̳n̳g̳e̳n̳i̳e̳r̳í̳a̳ i̳n̳f̳o̳r̳m̳á̳t̳i̳c̳a̳ c̳u̳a̳l̳ e̳s̳ l̳a̳ d̳u̳r̳a̳c̳i̳ó̳n̳ d̳e̳ d̳i̳c̳h̳a̳ c̳a̳r̳r̳e̳r̳a̳ ?̳ ",
			options:       []string{"2. ", "4. ", "6. "},
			correctOption: 3},
		{
			questionStr:   "D̳e̳ a̳c̳u̳e̳r̳d̳o̳ c̳o̳n̳ l̳a̳s̳ e̳s̳t̳a̳d̳í̳s̳t̳i̳c̳a̳s̳ d̳e̳ l̳a̳ f̳a̳c̳u̳l̳t̳a̳d̳,̳ e̳n̳ c̳u̳a̳n̳t̳o̳ t̳i̳e̳m̳p̳o̳ l̳o̳g̳r̳a̳ r̳e̳c̳i̳b̳i̳r̳s̳e̳ e̳l̳ e̳s̳t̳u̳d̳i̳a̳n̳t̳e̳ p̳r̳o̳m̳e̳d̳i̳o̳ d̳e̳ i̳n̳g̳e̳n̳i̳e̳r̳í̳a̳ i̳n̳f̳o̳r̳m̳á̳t̳i̳c̳a̳?̳?̳",
			options:       []string{"4. ", "6. ", "9. "},
			correctOption: 3},
		{
			questionStr:   " Q̳u̳e̳ e̳s̳ e̳l̳ S̳I̳U̳ g̳u̳a̳r̳a̳n̳i̳  ?̳",
			options:       []string{"El sistema de autogestion de la facultad. ", "Un aula de la facultad. ", "Una lengua indigena. "},
			correctOption: 1},
		{
			questionStr:   "G̳e̳n̳e̳r̳a̳l̳m̳e̳n̳t̳e̳ c̳u̳a̳n̳d̳o̳ s̳o̳n̳ l̳a̳s̳ i̳n̳s̳c̳r̳i̳p̳c̳i̳o̳n̳e̳s̳ a̳ l̳a̳s̳ m̳a̳t̳e̳r̳i̳a̳s̳ ?̳",
			options:       []string{"2 Semanas antes del comienzo de clases.  ", "1 Semana antes del comienzo de clases. ", "5 semanas antes del comienzo de clases. "},
			correctOption: 2},
		{
			questionStr:   "S̳i̳ a̳p̳r̳u̳e̳b̳o̳ e̳l̳ f̳i̳n̳a̳l̳ d̳e̳ u̳n̳a̳ m̳a̳t̳e̳r̳i̳a̳ d̳o̳n̳d̳e̳ p̳u̳e̳d̳o̳ e̳n̳c̳o̳n̳t̳r̳a̳r̳ l̳a̳ n̳o̳t̳a̳ e̳n̳ e̳l̳ s̳i̳s̳t̳e̳m̳a̳ ?̳ ",
			options:       []string{"En inscripciones a materias. ", "En inscripciones a exámenes. ", "En actuación provisoria. "},
			correctOption: 3},
		{
			questionStr:   "C̳u̳a̳l̳ e̳s̳ l̳a̳ c̳a̳t̳e̳d̳r̳a̳ p̳r̳e̳f̳e̳r̳i̳d̳a̳ p̳o̳r̳ l̳o̳s̳ e̳s̳t̳u̳d̳i̳a̳n̳t̳e̳s̳ a̳l̳ m̳o̳m̳e̳n̳t̳o̳ d̳e̳ c̳u̳r̳s̳a̳r̳ a̳l̳g̳e̳b̳r̳a̳ I̳I̳?̳",
			options:       []string{"Grynberg. ", "Lopez. ", "Alvarez. "},
			correctOption: 1},
		{
			questionStr:   "S̳e̳g̳u̳n̳ e̳l̳ p̳l̳a̳n̳ d̳e̳ e̳s̳t̳u̳d̳i̳o̳s̳ d̳e̳ I̳n̳g̳ I̳n̳f̳o̳r̳m̳a̳t̳i̳c̳a̳,̳ c̳u̳a̳l̳ d̳e̳ l̳a̳s̳ s̳i̳g̳u̳i̳e̳n̳t̳e̳s̳ m̳a̳t̳e̳r̳i̳a̳s̳ n̳o̳ e̳s̳ c̳o̳r̳r̳e̳l̳a̳t̳i̳v̳a̳ d̳e̳ l̳a̳ m̳a̳t̳e̳r̳i̳a̳ T̳e̳o̳r̳i̳a̳ d̳e̳ ̳l̳e̳n̳g̳u̳a̳j̳e̳ ?̳ ̳ ",
			options:       []string{"Algoritmos y Programación II. ", "Algoritmos y Programación I. ", "Fisica II. "},
			correctOption: 1},
		{
			questionStr:   "Q̳u̳e̳ p̳a̳s̳a̳ s̳i̳ m̳e̳ a̳n̳o̳t̳o̳ e̳n̳ u̳n̳a̳ m̳a̳t̳e̳r̳i̳a̳ y̳ l̳u̳e̳g̳o̳ d̳e̳c̳i̳d̳o̳ n̳o̳ c̳u̳r̳s̳a̳r̳l̳a̳?̳ ",
			options:       []string{"Te afecta el promedio. ", "Por 2 cuartrimestres no te podes volver a anotar. ", "Nada. "},
			correctOption: 3},
	}
	return questions
}
