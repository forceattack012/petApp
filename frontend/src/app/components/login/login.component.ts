import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  s = '';
  loginForm = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  })

  constructor(private userService: UserService, private activatedRoute: ActivatedRoute) { }

  ngOnInit(): void {
    this.s = this.activatedRoute.snapshot.paramMap.get('s')
  }

  Login() {
    if (this.loginForm.invalid) {
      return;
    }
    this.userService.login(this.loginForm.value).subscribe(result => {
      this.userService.saveToken(result.token);
      this.userService.saveName(result.name);
      window.location.href = "/"
    })
  }
}
