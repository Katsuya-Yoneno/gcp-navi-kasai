export class Question {
    id: string;
    question: string = '';
    ans1: string = '';
    ans2: string = '';
    ans3: string = '';
    ans4: string = '';
    correct_ans: string = '';
  
    constructor(id: string, question: string, ans1: string, ans2: string, ans3: string, ans4: string, correct_ans: string) {
      this.id = id;
      this.question = question;
      this.ans1 = ans1;
      this.ans2 = ans2;
      this.ans3 = ans3;
      this.ans4 = ans4;
      this.correct_ans = correct_ans;
    }
  }