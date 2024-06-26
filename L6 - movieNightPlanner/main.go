package main

import "fmt"

func main() {
	// Initialize slices to represent the movie collection
	titles := []string{}
	genres := []string{}
	durations := []int{}

	var choice int

	for {

		fmt.Println("Movie Night Planner")
		fmt.Println("1. Add a new movie")
		fmt.Println("2. Remove a movie by title")
		fmt.Println("3. List all movies of a specific genre")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add a new movie
			var title, genre string
			var duration int
			fmt.Print("Enter movie title: ")
			fmt.Scan(&title)
			fmt.Print("Enter movie genre: ")
			fmt.Scan(&genre)
			fmt.Print("Enter movie duration in minutes: ")
			fmt.Scan(&duration)

			titles = append(titles, title)
			genres = append(genres, genre)
			durations = append(durations, duration)

		case 2:
			// Remove a movie by title
			var removeTitle string
			fmt.Print("Enter the title of the movie to remove: ")
			fmt.Scan(&removeTitle)

			for i := 0; i < len(titles); i++ {
				if titles[i] == removeTitle {
					titles = append(titles[:i], titles[i+1:]...)
					genres = append(genres[:i], genres[i+1:]...)
					durations = append(durations[:i], durations[i+1:]...)
					break
				}
			
			}

		case 3:
			// List all movies of a specific genre
			var listGenre string
			fmt.Print("Enter the genre to list movies: ")
			fmt.Scan(&listGenre)

			fmt.Println("Movies of genre:", listGenre)
			for i := 0; i < len(genres); i++ {
				if genres[i] == listGenre {
					fmt.Printf("Title: %s, Duration: %d minutes\n", titles[i], durations[i])
				}
			}

		case 4:
			// Exit
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
