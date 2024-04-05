import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EmailPendingComponent } from './email-pending.component';

describe('EmailPendingComponent', () => {
  let component: EmailPendingComponent;
  let fixture: ComponentFixture<EmailPendingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EmailPendingComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(EmailPendingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
