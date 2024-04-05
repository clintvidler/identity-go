import { Component } from '@angular/core';
import { IdentityService } from '../../../../services/identity.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-email-pending',
  standalone: true,
  imports: [],
  templateUrl: './email-pending.component.html',
  styleUrl: './email-pending.component.scss',
})
export class EmailPendingComponent {
  key: string | null;

  constructor(
    private route: ActivatedRoute,
    private identityService: IdentityService
  ) {
    this.key = this.route.snapshot.paramMap.get('key') || '';

    this.identityService.finishUpdateEmail(this.key).subscribe((res) => {
      this.identityService.logout().subscribe();
    });
  }
}
