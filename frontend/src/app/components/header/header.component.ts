import { Component, OnInit } from '@angular/core';
import { NgbNavConfig } from '@ng-bootstrap/ng-bootstrap';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  name: string;

  constructor(config: NgbNavConfig, private userService: UserService) {
    config.destroyOnHide = false;
		config.roles = false;
  }

  ngOnInit(): void {
    this.name = this.userService.getName();
  }

}
