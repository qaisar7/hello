package main

import (
	"fmt"

	"github.com/qaisar7/strutil"
)

type Tasks struct {
	impact   int
	deadline int
	//job string
}

func AnalyzeGrades(tsks []Tasks, totalDays int) int {
	gradeImprovement := 0
	for i, j := 0, 0; i < totalDays && j < totalDays; i, j = i+1, j+1 {
		fmt.Printf ("Day:%d For task: %v , Days left:%d\n",i,tsks[j],tsks[j].deadline-i)
		if tsks[j].deadline-i > 0 {
			gradeImprovement = gradeImprovement + tsks[j].impact
			fmt.Printf("Adding task %v \n", tsks[j])
		} else  {
			for ;tsks[j].deadline - i <=0 && j < totalDays;  {
				fmt.Printf("Skipping task %v\n", tsks[j])
				j=j+1
				if j >= totalDays { break }
			}
			if j >= totalDays { break }
			if tsks[j].deadline-i > 0 {
				fmt.Printf ("Day:%d For task: %v , Days left:%d\n",i,tsks[j],tsks[j].deadline-i)
				gradeImprovement = gradeImprovement + tsks[j].impact
				fmt.Printf("Adding task %v \n", tsks[j])
			}
		}
		fmt.Printf("GradeImprovment: %d \n", gradeImprovement)
	}
	return gradeImprovement
}

func main() {

	fmt.Println("The tasks are shown as {Impact, Deadline} in order from left to right in '[]'")

//	tsk := []Tasks{{4, 3}, {13, 1}, {10, 2}, {10, 5}, {15, 1}, {15, 2}}
//	tsk := []Tasks{{70, 4}, {60, 2}, {50, 4}, {40, 3}, {30, 1}, {20, 4}, {10,6}}
//	tsk := []Tasks{{6, 1}, {7, 1}, {2, 3}, {1, 3}, {4, 2}, {5, 2}, {6, 1}}
	tsk := []Tasks{{20,4}, {10,1}, {40,1},{30,1}}
//	tsk := []Tasks{{100,2,"a"}, {19,1, "b"}, {27,2, "c"}, {25,1,"d"}, {15,3,"e"}}
	fmt.Println("Raw input ", tsk, "\n")
	currentGrade := 0
	totalDays := len(tsk)

	currentGrade = currentGrade + AnalyzeGrades(tsk, totalDays)
	fmt.Println("Grade improvement according to raw input is ", currentGrade, "\n\n")
	
	currentGrade = 0
	for i := 0; i < len(tsk); i++ {
		if tsk[i].deadline - i > 0 {
			for j := i + 1; j < len(tsk); j++ {
					fmt.Printf("Comparing key: %v with: %v \n", tsk[i], tsk[j])
					if tsk[j].impact > tsk[i].impact && tsk[j].deadline-i >0 {
						fmt.Printf("Flipping key: %v with: %v \n", tsk[i], tsk[j])
						tsk[i], tsk[j] = tsk[j], tsk[i]
					} else if tsk[j].impact == tsk[i].impact {
						if tsk[j].deadline < tsk[i].deadline && tsk[j].deadline-i >0 {
							fmt.Printf("Equal impact, flipping key: %v with: %v \n", tsk[i], tsk[j])
							tsk[i], tsk[j] = tsk[j], tsk[i]							
						}
					}
				} 
		} else {  //Deadline for this task has expired. Look for the next biggest impact and swap this one with it
			for j:= i+1; j<len(tsk); j++ {
				if tsk[j].deadline - i > 0 {
					fmt.Printf("Deadline expired for %v. Flipping with: %v \n", tsk[i], tsk[j])
					tsk[i], tsk[j] = tsk[j], tsk[i]
					i = i-1
					break
				}
				
			}			
		}
	}
	fmt.Println("Sorted tasks ", tsk, "\n")
	currentGrade = currentGrade + AnalyzeGrades(tsk, totalDays)
	fmt.Println("Grade improvement according to sorted tasks is ", currentGrade)
	fmt.Println(strutil.Reverse("wow man"))
}
