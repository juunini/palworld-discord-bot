import { randomID } from '../lib/randomID.js';

class Input extends HTMLElement {
  static get observedAttributes() {
    return ['value'];
  }

  inputID = 'input-' + randomID();
  div = document.createElement('div');
  input = document.createElement('input');
  label = document.createElement('label');
  tooltip = document.createElement('button');
  value = false;

  constructor() { super(); }

  connectedCallback() {
    this.id = this.getAttribute('id');

    this.appendChild(this.div);

    this._decorateDiv();
    this._decorateLabel();
    this._decorateInput();

    this.input.addEventListener('change', (e) => {
      this.value = e.target.value;
    });
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'value') {
      this.value = newValue;
      this.input.value = this.value;
    }
  }

  _decorateDiv() {
    this.div.className = `input-group ${this.getAttribute('class') || ''}`;
    this.div.appendChild(this.label);
    this.div.appendChild(this.input);
  }

  _decorateLabel() {
    this.label.className = 'input-group-text';
    this.label.htmlFor = this.inputID;
    this.label.textContent = this.getAttribute('label');

    if (!this.getAttribute('tooltip')) {
      return;
    }

    this.label.appendChild(this.tooltip);
    this._decorateTooltip();
  }

  _decorateInput() {
    this.input.classList.add('form-control');
    this.input.type = this.getAttribute('type');
    this.input.id = this.inputID;
    this.input.placeholder = this.getAttribute('placeholder') || '';
  }

  _decorateTooltip() {
    this.tooltip.className = 'btn btn-outline-secondary ms-2 d-inline-flex flex-row justify-content-center align-items-center p-1';
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
}

customElements.define('custom-input', Input);
