import { randomID } from '../lib/randomID.js';

class Switch extends HTMLElement {
  static get observedAttributes() {
    return ['checked'];
  }

  inputID = 'switch-' + randomID();
  div = document.createElement('div');
  checkbox = document.createElement('input');
  label = document.createElement('label');
  tooltip = document.createElement('button');
  checked = false;

  turnOffPrompt = '';
  modalOpenButton = document.createElement('button');
  close = 'Close';
  confirm = 'Confirm';

  constructor() { super(); }

  connectedCallback() {
    this.id = this.getAttribute('id');

    this.appendChild(this.div);

    this._decorateDiv();
    this._decorateInput();
    this._decorateLabel();

    this.checkbox.addEventListener('change', (e) => {
      if (!e.target.checked && this.turnOffPrompt) {
        this.checkbox.checked = true;
        this.modalOpenButton.click();
        return;
      }

      this.checked = e.target.checked;
      this.checkbox.checked = this.checked;
    });

    this.turnOffPrompt = this.getAttribute('turn-off-prompt') || '';

    if (this.turnOffPrompt) {
      this._addModal();
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'checked') {
      this.checked = newValue === 'true' || newValue === true;
      this.checkbox.checked = this.checked;
    }
  }

  _decorateDiv() {
    this.div.className = `form-check form-switch ${this.getAttribute('class') || ''}`;
    this.div.appendChild(this.checkbox);
    this.div.appendChild(this.label);
  }

  _decorateInput() {
    this.checkbox.classList.add('form-check-input');
    this.checkbox.type = 'checkbox';
    this.role = 'switch';
    this.checkbox.id = this.inputID;
  }

  _decorateLabel() {
    this.label.classList.add('form-check-label');
    this.label.htmlFor = this.inputID;
    this.label.textContent = this.getAttribute('label') || '';

    if (!this.getAttribute('tooltip')) {
      return;
    }

    this.label.className = 'form-check-label d-flex flex-row justify-content-start align-items-center';
    this.label.appendChild(this.tooltip);
    this._decorateTooltip();
  }

  _decorateTooltip() {
    this.tooltip.className = 'btn btn-outline-secondary ms-2 d-flex flex-row justify-content-center align-items-center p-1';
    this.tooltip.type = 'button';
    this.tooltip.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-question-circle" viewBox="0 0 16 16">
      <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16"/>
      <path d="M5.255 5.786a.237.237 0 0 0 .241.247h.825c.138 0 .248-.113.266-.25.09-.656.54-1.134 1.342-1.134.686 0 1.314.343 1.314 1.168 0 .635-.374.927-.965 1.371-.673.489-1.206 1.06-1.168 1.987l.003.217a.25.25 0 0 0 .25.246h.811a.25.25 0 0 0 .25-.25v-.105c0-.718.273-.927 1.01-1.486.609-.463 1.244-.977 1.244-2.056 0-1.511-1.276-2.241-2.673-2.241-1.267 0-2.655.59-2.75 2.286m1.557 5.763c0 .533.425.927 1.01.927.609 0 1.028-.394 1.028-.927 0-.552-.42-.94-1.029-.94-.584 0-1.009.388-1.009.94"/>
    </svg>`;
    this.tooltip.setAttribute('data-bs-toggle', 'tooltip');
    this.tooltip.setAttribute('data-bs-placement', 'top');
    this.tooltip.setAttribute('title', this.getAttribute('tooltip'));
    this.tooltip.style.border = 'none';
    this.tooltip.style.borderRadius = '50%';
  }

  _addModal() {
    this.modalOpenButton.type = 'button';
    this.modalOpenButton.style.display = 'none';
    this.modalOpenButton.setAttribute('data-bs-toggle', 'modal');
    this.modalOpenButton.setAttribute('data-bs-target', '#' + this.inputID + '-modal');
    
    const modal = document.createElement('div');
    modal.className = 'modal fade';
    modal.id = this.inputID + '-modal';
    modal.setAttribute('tabindex', '-1');
    modal.setAttribute('aria-labelledby', this.inputID + '-modal-label');
    modal.setAttribute('aria-hidden', 'true');
    
    const modalDialog = document.createElement('div');
    modalDialog.className = 'modal-dialog';

    const modalContent = document.createElement('div');
    modalContent.className = 'modal-content';

    const modalBody = document.createElement('div');
    modalBody.className = 'modal-body';
    modalBody.textContent = this.turnOffPrompt;

    const modalFooter = document.createElement('div');
    modalFooter.className = 'modal-footer';

    const modalClose = document.createElement('button');
    modalClose.type = 'button';
    modalClose.className = 'btn btn-secondary';
    modalClose.setAttribute('data-bs-dismiss', 'modal');
    modalClose.textContent = this.getAttribute('close') || 'Close';

    const modalConfirm = document.createElement('button');
    modalConfirm.type = 'button';
    modalConfirm.className = 'btn btn-primary';
    modalConfirm.textContent = this.getAttribute('confirm') || 'Confirm';
    modalConfirm.addEventListener('click', () => {
      this.checked = false;
      this.checkbox.checked = false;
      modalClose.click();
    });

    modalFooter.appendChild(modalClose);
    modalFooter.appendChild(modalConfirm);
    modalContent.appendChild(modalBody);
    modalContent.appendChild(modalFooter);
    modalDialog.appendChild(modalContent);
    modal.appendChild(modalDialog);
    this.appendChild(modal);
    this.appendChild(this.modalOpenButton);
  }
}

customElements.define('custom-switch', Switch);
