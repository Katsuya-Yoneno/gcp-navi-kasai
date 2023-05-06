export class Question {
    id: number;
    question: string = '';
    category: string = '';
    ans1: string = '';
    ans2: string = '';
    ans3: string = '';
    ans4: string = '';
    correct_ans: number = 0;
  
    constructor(id: number, question: string, category: string, ans1: string, ans2: string, ans3: string, ans4: string, correct_ans: number) {
      this.id = id;
      this.question = question;
      this.category = category;
      this.ans1 = ans1;
      this.ans2 = ans2;
      this.ans3 = ans3;
      this.ans4 = ans4;
      this.correct_ans = correct_ans;
    }
  }