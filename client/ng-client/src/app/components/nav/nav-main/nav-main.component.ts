import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-nav-main',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './nav-main.component.html',
  styleUrl: './nav-main.component.scss',
})
export class NavMainComponent {
  showMenu = false;

  toggleNavbar() {
    this.showMenu = !this.showMenu;
  }
}
