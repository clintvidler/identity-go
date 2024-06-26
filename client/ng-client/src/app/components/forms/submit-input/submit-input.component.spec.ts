import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubmitInputComponent } from './submit-input.component';

describe('SubmitInputComponent', () => {
  let component: SubmitInputComponent;
  let fixture: ComponentFixture<SubmitInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SubmitInputComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SubmitInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
