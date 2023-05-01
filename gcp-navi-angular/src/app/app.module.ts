import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { QuestionListComponent } from './questions/questions.component';
import { HttpClientModule } from '@angular/common/http';
import { QuestionService } from './questions.service';
import { FirstOptionsComponent } from './first-options/first-options.component';
import { ChallengeQuestionsComponent } from './challenge-questions/challenge-questions.component';
import { NgxPaginationModule } from 'ngx-pagination';


@NgModule({
  declarations: [
    AppComponent,
    QuestionListComponent,
    FirstOptionsComponent,
    ChallengeQuestionsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    NgxPaginationModule
  ],
  providers: [QuestionService],
  bootstrap: [AppComponent]
})
export class AppModule { }
