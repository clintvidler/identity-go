import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { IdentityService } from '../../../services/identity.service';
import { User } from '../../../interfaces/user';

@Component({
  selector: 'app-nav-main',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './nav-main.component.html',
  styleUrl: './nav-main.component.scss',
})
export class NavMainComponent implements OnInit {
  showMenu = false;
  user: User | null = null;

  constructor(private identityService: IdentityService) {}

  ngOnInit(): void {
    this.identityService.currentUserSubject.subscribe((user: User | null) => {
      this.user = user;
    });
    // this.user = this.identityService.currentUser;
  }

  toggleNavbar() {
    this.showMenu = !this.showMenu;
  }

  authItems: MenuItem[] = [
    {
      text: 'profile',
      link: '/user',
    },
    {
      text: 'logout',
      link: '/logout',
    },
  ];

  unAuthItems: MenuItem[] = [
    {
      text: 'login',
      link: '/login',
    },
    {
      text: 'register',
      link: '/register',
    },
  ];
}

interface MenuItem {
  text: string;
  link: string;
}
