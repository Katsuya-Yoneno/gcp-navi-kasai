import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QuestionListComponent } from './questions/questions.component';
import { FirstOptionsComponent } from './first-options/first-options.component';

const routes: Routes = [
  { path: '', component: FirstOptionsComponent },
  { path: 'questions', component: QuestionListComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
