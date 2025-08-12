import { Component, signal } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import {
  CdkDragDrop,
  moveItemInArray,
  transferArrayItem,
  CdkDrag,
  CdkDropList,
} from '@angular/cdk/drag-drop';
import { OidcSecurityService } from 'angular-auth-oidc-client'
import { Observable } from 'rxjs';
import { AppService } from './app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.html',
  styleUrl: './app.css',
  imports: [FormsModule, CdkDropList, CdkDrag]
})
export class App {

  protected readonly title = signal('angularjs');
  configuration$: Observable<any>;
  userData$: Observable<any>;

  isAuthenticated = false;
  showBoard = false;
  statusButtonSave = false;

  task: string = '';

  alertStatus = {
    visible: false,
    class: '',
    label: ''
  };

  boardId = undefined;
  todo: Array<string> = [];
  doing: Array<string> = [];
  done: Array<string> = [];

  constructor(private oidcSecurityService: OidcSecurityService, private appService: AppService) {
    this.configuration$ = this.oidcSecurityService.getConfiguration();
    this.userData$ = this.oidcSecurityService.userData$;
  }

  ngOnInit(): void {
    this.oidcSecurityService.checkAuth().subscribe(() => {
      this.oidcSecurityService.isAuthenticated$.subscribe(({ isAuthenticated }) => {
        this.isAuthenticated = isAuthenticated;

        if (this.isAuthenticated) {
          this.oidcSecurityService.getAccessToken().subscribe((token) => {
            if (token) {
              window.localStorage.setItem('token', token);
              this.getBoards();
            }
          });
        }
      });
    });
  }

  login(): void {
    this.oidcSecurityService.authorize();
  }

  logout(): void {
    if (window.sessionStorage) {
      window.sessionStorage.clear();
      localStorage.clear();
    }

    this.oidcSecurityService.logoffAndRevokeTokens().subscribe(() => {
      window.location.href = process.env['APP_AWS_REDIRECT_URL'] || ''
    });
  }

  addTask(): void {
    if (this.task.trim().length === 0)
      return;

    this.todo.push(this.task);
    this.task = '';
  }

  removeTask(idx: number, step: string): void {
    switch (step) {
      case 'todo': {
        this.todo.splice(idx, 1);
        break;
      }
      case 'doing': {
        this.doing.splice(idx, 1);
        break;
      }
      default: {
        this.done.splice(idx, 1);
      }
    }
  }

  getBoards(): void {
    this.appService.get()
      .subscribe({
        next: (boards) => {
          if (boards && boards.length > 0) {
            this.boardId = boards[0].board_id;
            const tasks = JSON.parse(boards[0].tasks);
            this.todo = tasks.todo
            this.doing = tasks.doing
            this.done = tasks.done
          }
        },
        error: (err: HttpErrorResponse) => {
          this.handleMessage('alert alert-danger', err.error);
        },
        complete: () => {
          this.showBoard = true;
        }
      });
  }

  saveTasks(): void {
    const board = {
      "boardid": this.boardId,
      "todo": this.todo,
      "doing": this.doing,
      "done": this.done
    }

    this.statusButtonSave = true;
    this.appService.save(board)
      .subscribe({
        next: (result) => {
          this.boardId = result.boardid;
          this.handleMessage('alert alert-success', result.msg);
        },
        error: (err: HttpErrorResponse) => {
          this.handleMessage('alert alert-danger', err.error);
        },
        complete: () => {
          this.statusButtonSave = false;
        }
      });
  }

  handleMessage(alertClass: string, alertClassLabel: string): void {
    this.alertStatus.visible = true;
    this.alertStatus.class = alertClass;
    this.alertStatus.label = alertClassLabel;
  }

  drop(event: CdkDragDrop<string[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(
        event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex,
      );
    }
  }
}
