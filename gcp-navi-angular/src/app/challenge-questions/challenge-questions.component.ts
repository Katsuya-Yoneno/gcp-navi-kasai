import { Component, OnInit, ViewChild } from '@angular/core';
import { Question } from '../question';
import { QuestionService } from '../questions.service';
import { PaginationControlsComponent } from 'ngx-pagination';

@Component({
  selector: 'app-questions',
  templateUrl: './challenge-questions.component.html',
  styleUrls: ['./challenge-questions.component.css']
})
export class ChallengeQuestionsComponent implements OnInit {
  @ViewChild(PaginationControlsComponent, { static: false }) paginator!: PaginationControlsComponent;
  p: number = 1; 
  isCorrect: boolean = false;
  isAnswered: boolean = false;
  questions: Question[] = [];

  constructor(private questionService: QuestionService) { }

  ngOnInit(): void {
    this.questionService.getQuestions().subscribe((questions: Question[]) => {
      this.questions = questions;
      console.log('問題一覧', this.questions);
    });
  }
  
  //問題の正解・不正解を判定
  showAnswer(selectedAnswer: number, correctAnswer: number) {
    console.log('answer:', correctAnswer)
    this.isAnswered = true;
    if (selectedAnswer === correctAnswer) {
      this.isCorrect = true;
    } else {
      this.isCorrect = false;
    }
  }

  // ページが切り替わったときに、情報をリセットする
  paginationControlsEvent(pageNumber: number): void {
    this.p = pageNumber;
    this.isAnswered = false;
    this.isCorrect = false;
  }
  
}