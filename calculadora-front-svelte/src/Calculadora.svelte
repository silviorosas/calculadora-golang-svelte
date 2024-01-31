
<script>
  import { navigate } from "svelte-navigator";

  

    let displayValue = '';
    let errorMessage = '';
  
    async function addToDisplay(value) {
  if (value === '=') {
    let expression = displayValue.trim();
    let number1 = '';
    let operator = '';
    let number2 = '';
  
    // Manejar el caso de número negativo
    if (expression.startsWith('-')) {
      number1 += '-';
      expression = expression.slice(1).trim(); // Eliminar el signo negativo del comienzo
    }
  
    // Recorrer la expresión para encontrar el operador
    for (let i = 0; i < expression.length; i++) {
      if (expression[i] === '+' || expression[i] === '-' || expression[i] === '*' || expression[i] === '/') {
        operator = expression[i];
        number2 = expression.slice(i + 1).trim();
        break;
      } else {
        number1 += expression[i];
      }
    }
  
    // Si no se encontró un operador, mostrar un mensaje de error
    if (!operator) {
      errorMessage = 'Expresión inválida';
      return;
    }
  
    // Agrega console.log para verificar los valores
    console.log("Número 1:", number1);
        console.log("Operador:", operator);
        console.log("Número 2:", number2);
  
    try {
      const response = await fetch('http://localhost:9000/calculate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
          number1: parseFloat(number1), 
          operator, 
          number2: parseFloat(number2) 
        })
      });
  
      const data = await response.json();
      if (data.error) {
        errorMessage = data.error;
        displayValue = '';
      } else {
        displayValue = data.result.toString();
        errorMessage = '';
      }
    } catch (error) {
      console.error('Error:', error);
    }
  } else {
    displayValue += value.toString();
  }
  }
  
    function clearDisplay() {
      displayValue = '';
      errorMessage = '';
    }
  
    document.addEventListener('keydown', function(event) {
      const key = event.key;
      // @ts-ignore
      if (!isNaN(key) || ['+', '-', '*', '/', '.', 'Enter', 'Backspace'].includes(key)) {
        if (key === 'Enter') {
          addToDisplay('=');
        } else if (key === 'Backspace') {
          clearDisplay();
        } else {
          addToDisplay(key);
        }
      }
    });
  
    function deleteLastDigit() {
    displayValue = displayValue.slice(0, -1);

    function goToHistory() {
    navigate("/historial");
  }
  }
  </script>
  
  <h2>Calculadora Go + Svelte  by Silvio Rosas</h2>

  
  <div class="calculator">
    <div class="display">
      {errorMessage ? errorMessage : displayValue}
    </div>
    <div class="buttons">
      <button class="button" on:click={() => addToDisplay('7')}>7</button>
      <button class="button" on:click={() => addToDisplay('8')}>8</button>
      <button class="button" on:click={() => addToDisplay('9')}>9</button>
      <button class="button" on:click={clearDisplay}>C</button>
      <button class="button" on:click={() => addToDisplay('4')}>4</button>
      <button class="button" on:click={() => addToDisplay('5')}>5</button>
      <button class="button" on:click={() => addToDisplay('6')}>6</button>
      <button class="button" on:click={() => addToDisplay('*')}>*</button>
      <button class="button" on:click={() => addToDisplay('1')}>1</button>
      <button class="button" on:click={() => addToDisplay('2')}>2</button>
      <button class="button" on:click={() => addToDisplay('3')}>3</button>
      <button class="button" on:click={() => addToDisplay('-')}>-</button>
      <button class="button" on:click={() => addToDisplay('0')}>0</button>
      <button class="button" on:click={() => addToDisplay('.')}>.</button>
      <button class="button" on:click={() => addToDisplay('+')}>+</button>     
      <button class="button" on:click={() => addToDisplay('/')}>/</button>
      <button class="button" on:click={deleteLastDigit}>←</button>
      <button class="button" on:click={() => addToDisplay('=')}>=</button>
      <button class="button" on:click={() => navigate("/historial")}>m</button>
      
    </div>
  </div>
  
  
  
  
  
  
  
  <style>
    /* Estilos CSS para la calculadora */
    .calculator {
      width: 240px;
      background-color: #92bfcc;
      border: 1px solid #1d7a96;
      border-radius: 5px;
      padding: 10px;
      margin: 0 auto;
    }
  
    .display {
    width: calc(100% - 8px); /* Ajustar el ancho para considerar el borde */
    border: 1px solid #1d7a96;
    border-radius: 5px;
    height: 40px;
    font-size: 20px;
    font-weight: 500;
    text-align: right;
    margin-bottom: 10px;
    padding: 4px;
    margin: 0 auto; /* Centrar horizontalmente */
    margin-bottom: 10px;
  }
  
  
   
  
    .buttons {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: 5px;
    }
  
    .button {
      background-color: #2b829c;
      border: 1px solid #131212;
      border-radius: 5px;
      padding: 10px;
      cursor: pointer;
      transition: background-color 0.3s ease;
    } 
  
    .button:hover {
      background-color: #ccc;
    }
  </style>
  
  