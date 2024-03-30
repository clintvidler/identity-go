import { Directive, Inject, Injector, OnInit } from '@angular/core';
import {
  ControlValueAccessor,
  FormControl,
  FormControlDirective,
  FormControlName,
  FormGroupDirective,
  NgControl,
  Validators,
} from '@angular/forms';
import { Subject, distinctUntilChanged, startWith, takeUntil, tap } from 'rxjs';

@Directive({
  selector: '[appControlValueAccessor]',
  standalone: true,
})
export class ControlValueAccessorDirective<T>
  implements ControlValueAccessor, OnInit
{
  control: FormControl = new FormControl();
  isRequired: boolean = false;

  private _isDisabled = false;
  private _destroy$ = new Subject<void>();
  private _onTouched!: () => T;

  constructor(@Inject(Injector) private injector: Injector) {}

  ngOnInit() {
    this.setFormControl();
    this.isRequired = this.control?.hasValidator(Validators.required) ?? false;
  }

  setFormControl() {
    try {
      const formControl = this.injector.get(NgControl);

      switch (formControl.constructor) {
        case FormControlName:
          this.control = this.injector
            .get(FormGroupDirective)
            .getControl(formControl as FormControlName);
          break;
        default:
          this.control = (formControl as FormControlDirective)
            .form as FormControl;
      }
    } catch (err) {
      this.control = new FormControl();
    }
  }

  writeValue(value: T): void {
    this.control
      ? this.control.setValue(value)
      : (this.control = new FormControl(value));
  }

  registerOnChange(fn: (value: T | null) => T): void {
    this.control?.valueChanges
      .pipe(
        takeUntil(this._destroy$),
        startWith(this.control.value),
        distinctUntilChanged(),
        tap((value) => fn(value))
      )
      .subscribe(() => this.control?.markAllAsTouched());
  }

  registerOnTouched(fn: any): void {
    this._onTouched = fn;
  }

  setDisabledState?(isDisabled: boolean): void {
    this._isDisabled = isDisabled;
  }
}
