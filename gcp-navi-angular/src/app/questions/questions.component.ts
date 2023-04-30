import { Component, OnInit } from '@angular/core';
import { Question } from '../question';
import { QuestionService } from '../questions.service';

@Component({
  selector: 'app-questions',
  templateUrl: './questions.component.html',
  styleUrls: ['./questions.component.css']
})
export class QuestionListComponent implements OnInit {
  questions: Question[] = [];

  constructor(private questionService: QuestionService) { }

  ngOnInit(): void {
    this.questionService.getQuestions().subscribe((questions: Question[]) => {
      this.questions = questions;
      console.log('問題一覧', this.questions);
    });
  }
}
