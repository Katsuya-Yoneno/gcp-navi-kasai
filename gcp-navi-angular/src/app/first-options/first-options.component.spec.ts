import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FirstOptionsComponent } from './first-options.component';

describe('FirstOptionsComponent', () => {
  let component: FirstOptionsComponent;
  let fixture: ComponentFixture<FirstOptionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FirstOptionsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FirstOptionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
