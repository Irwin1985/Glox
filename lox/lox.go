package lox

import (
	"Glox/scanner"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// hadError will ensure don't try to execute code that has a known error.
var HadError = true

func RunFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("file does not exist: %s", path)
		os.Exit(65)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	run(string(bytes))
	if HadError {
		os.Exit(65)
	}
}

func RunPrompt() {
	mode := "debug"
	if mode != "debug" {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanned := scanner.Scan()
			if !scanned {
				return
			}
			fmt.Print("> ")
			line := scanner.Text()
			if len(line) <= 0 {
				continue
			}
			run(line)
			// we reset hadError if the user makes a mistake and
			// it shouldn't kill the entire session.
			HadError = false
		}
	} else {
		line := `
		// Your first Lox program!
		print "Hello, world!";

		true;  // Not false.
		false; // Not *not* false.

		1234;  // An integer.
		12.34; // A decimal number.
		
		"I am a string";
		"";    // The empty string.
		"123"; // This is a string, not a number.

		add + me;
		subtract - me;
		multiply * me;
		divide / me;
		
		-negateMe;

		less < than;
		lessThan <= orEqual;
		greater > than;
		greaterThan >= orEqual;
		
		1 == 2;         // false.
		"cat" != "dog"; // true.
		
		314 == "pi"; // false.
		
		123 == "123"; // false.

		!true;  // false.
		!false; // true.
		
		true and false; // false.
		true and true;  // true.
		
		false or false; // false.
		true or false;  // true.
		
		var average = (min + max) / 2;

		print "Hello, world!";

		"some expression";

		{
			print "One statement.";
			print "Two statements.";
		}
		  
		var imAVariable = "here is my value";
		var iAmNil;
		
		var breakfast = "bagels";
		print breakfast; // "bagels".
		breakfast = "beignets";
		print breakfast; // "beignets".
		
		if (condition) {
			print "yes";
		} else {
			print "no";
		}

		var a = 1;
		while (a < 10) {
		  print a;
		  a = a + 1;
		}
		
		for (var a = 1; a < 10; a = a + 1) {
			print a;
		}

		makeBreakfast(bacon, eggs, toast);

		makeBreakfast();

		fun printSum(a, b) {
			print a + b;
		}

		fun returnSum(a, b) {
			return a + b;
		}		

		fun addPair(a, b) {
			return a + b;
		}
		  
		fun identity(a) {
			return a;
		}
		  
		print identity(addPair)(1, 2); // Prints "3".		

		fun outerFunction() {
			fun localFunction() {
			  print "I'm local!";
			}
		  
			localFunction();
		}		

		fun returnFunction() {
			var outside = "outside";
		  
			fun inner() {
			  print outside;
			}
		  
			return inner;
		}
		  
		var fn = returnFunction();
		fn();

		class Breakfast {
			cook() {
			  print "Eggs a-fryin'!";
			}
		  
			serve(who) {
			  print "Enjoy your breakfast, " + who + ".";
			}
		}		

		// Store it in variables.
		var someVariable = Breakfast;
		
		// Pass it to functions.
		someFunction(Breakfast);
		
		var breakfast = Breakfast();
		print breakfast; // "Breakfast instance".
		
		breakfast.meat = "sausage";
		breakfast.bread = "sourdough";
		
		class Breakfast {
			serve(who) {
			  print "Enjoy your " + this.meat + " and " +
				  this.bread + ", " + who + ".";
			}
		  
			// ...
		}		

		class Breakfast {
			init(meat, bread) {
			  this.meat = meat;
			  this.bread = bread;
			}
		  
			// ...
		  }
		  
		var baconAndToast = Breakfast("bacon", "toast");
		baconAndToast.serve("Dear Reader");
		// "Enjoy your bacon and toast, Dear Reader."		

		class Brunch < Breakfast {
			drink() {
			  print "How about a Bloody Mary?";
			}
		}		

		var benedict = Brunch("ham", "English muffin");
		benedict.serve("Noble Reader");
		
		class Brunch < Breakfast {
			init(meat, bread, drink) {
			  super.init(meat, bread);
			  this.drink = drink;
			}
		}	
		fin			  
		`
		run(line)
	}
}

func run(source string) {
	s := scanner.NewScanner(source)
	tokens := s.ScanTokens()

	// For now, just print the tokens.
	for _, tok := range tokens {
		fmt.Println(tok.ToString())
	}
}
