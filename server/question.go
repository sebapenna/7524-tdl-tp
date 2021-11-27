package server

type Question struct {
	questionNumber int
	question       string
	options        []string
	correctOption  int
}

func CreateRandomQuestions() []Question {
	questions := []Question{
		{
			question:      "How many permanent teeth does a dog have?",
			options:       []string{"42", "40", "36"},
			correctOption: 1},
		{
			question:      "How many chukkers are there in a polo match?",
			options:       []string{"7", "6", "8"},
			correctOption: 2},
		{
			question:      "On average, how far away is the moon from the earth?",
			options:       []string{"384.400 km", "450.200 km", "285.700"},
			correctOption: 1},
		{
			question:      "Complete this Spice Girls lyric: “If you wanna be my [...], you gotta get with my friends”",
			options:       []string{"husband", "friend", "lover"},
			correctOption: 3},
		{
			question:      "In what decade was pop icon Madonna born?",
			options:       []string{"60s", "50s", "70s"},
			correctOption: 2},
		{
			question:      "Which European city hosted the 1936 Summer Olympics?",
			options:       []string{"Berlin", "Madrid", "London"},
			correctOption: 1},
		{
			question:      "Which country won the 1994 FIFA World Cup?",
			options:       []string{"France", "Germany", "Brazil"},
			correctOption: 3},
		{
			question:      "What is the capital city of Australia?",
			options:       []string{"Sydney", "Canberra", "Melbourne"},
			correctOption: 2},
		{
			question:      "Which US state was Donald Trump born in?",
			options:       []string{"Texas", "California", "New York"},
			correctOption: 3},
		{
			question:      "What is the capital of Finland?",
			options:       []string{"Helsinki", "Copenhagen", "Stockholm"},
			correctOption: 1},
		{
			question:      "What language is spoken in Brazil?",
			options:       []string{"Spanish", "Portuguese", "English"},
			correctOption: 2},
		{
			question:      "What temperature centigrade does water boil at?",
			options:       []string{"100", "80", "75"},
			correctOption: 1},
		{
			question:      "How many notes are there in a musical scale?",
			options:       []string{"7", "6", "5"},
			correctOption: 1},
		{
			question:      "What in the animal kingdom is a doe?",
			options:       []string{"A female donkey", "A female horse", "A female deer"},
			correctOption: 3},
		{
			question:      "What is the busiest airport in Britain called?",
			options:       []string{"Luton", "Heathrow", "Gatwick"},
			correctOption: 2},
		{
			question:      "Who is next in line to the British throne after Queen Elizabeth II?",
			options:       []string{"Prince Charles", "Prince Williams", "Prince Harry"},
			correctOption: 1},
		{
			question:      "What was the first single to be released by the band Oasis?",
			options:       []string{"Don’t Believe The Truth", "Dig Out Your Soul", "Supersonic"},
			correctOption: 3},
		{
			question:      "In which year did the original UK version of The Office air?",
			options:       []string{"2016", "2001", "2010"},
			correctOption: 2},
		{
			question:      "Glossectomy is the removal of all of or part of which body part?",
			options:       []string{"The tongue", "The lungs", "A finger"},
			correctOption: 1},
		{
			question:      "How many elements are in the periodic table?",
			options:       []string{"133", "125", "118"},
			correctOption: 3},
		{
			question:      "Which planet has the most moons?",
			options:       []string{"Jupiter", "Saturn", "Neptune"},
			correctOption: 2},
		{
			question:      "Where is the smallest bone in the human body located?",
			options:       []string{"Leg", "Nose", "Ear"},
			correctOption: 3},
		{
			question:      "In which year was the Nintendo 64 released in Europe?",
			options:       []string{"1997", "1995", "1999"},
			correctOption: 1},
		{
			question:      "Elon Musk is the CEO of which global brand.",
			options:       []string{"Facebook", "Amazon", "Tesla"},
			correctOption: 3},
		{
			question:      "In which year did Taylor Swift release her debut single, Love Story?",
			options:       []string{"2008", "2012", "2019"},
			correctOption: 1},
		{
			question:      "Who is the eldest Weasley sibling in Harry Potter?",
			options:       []string{"Bill Weasley", "Ron Weasly", "Fred Weasley"},
			correctOption: 1},
		{
			question:      "How many Pirates of the Caribbean films have been released?",
			options:       []string{"4", "5", "6"},
			correctOption: 2},
		{
			question:      "In Monsters Inc. what is Sulley’s full name?",
			options:       []string{"George P. Sullivan", "James P. Sullivan", "John P. Sullivan"},
			correctOption: 2},
		{
			question:      "Who discovered penicillin?",
			options:       []string{"Hermann Joseph Muller", "Marie Curie", "Alexander Fleming"},
			correctOption: 3},
		{
			question:      "Which year did the European Union first introduce the Euro as currency?",
			options:       []string{"1999", "2003", "2001"},
			correctOption: 1},
		{
			question:      "Which English Football League team holds the nickname The Cobblers?",
			options:       []string{"Chelsea", "Northampton Town", "Tottenhan Hotspurs"},
			correctOption: 2},
		{
			question:      "Who scored the fastest goal in Premier League history after just 7.69 seconds?",
			options:       []string{"Cristiano Ronaldo", "Shane Long", "Wayne Rooney"},
			correctOption: 2},
		{
			question:      "How many Grand Slam singles titles has Serena Williams won?",
			options:       []string{"23", "18", "29"},
			correctOption: 1},
		{
			question:      "How many Grand Slam titles has Andy Murray won?",
			options:       []string{"4", "5", "3"},
			correctOption: 3},
		{
			question:      "Where was Frida Kahlo born?",
			options:       []string{"El Salvador", "Costa Rica", "Mexico"},
			correctOption: 3},
		{
			question:      "What is the capital of Bulgaria?",
			options:       []string{"Bucarest", "Budapest", "Sofia"},
			correctOption: 3},
		{
			question:      "What is the official name of the clock tower commonly referred to as Big Ben?",
			options:       []string{"Elizabeth Tower", "Westminster Tower", "King's Cross Tower"},
			correctOption: 1}}
	return questions
}
