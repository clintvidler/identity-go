import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvalidKeyComponent } from './invalid-key.component';

describe('InvalidKeyComponent', () => {
  let component: InvalidKeyComponent;
  let fixture: ComponentFixture<InvalidKeyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InvalidKeyComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(InvalidKeyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
