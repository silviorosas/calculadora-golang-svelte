<script>

  import { navigate } from "svelte-navigator";

    let history = [];
  
    async function fetchHistory() {
      try {
        const response = await fetch('http://localhost:9000/history');
        if (!response.ok) {
          throw new Error('Error al obtener el historial de operaciones');
        }
        history = await response.json();
        formatDates(); // Formatear fechas al recibir el historial
      } catch (error) {
        console.error(error);
      }
    }
  
    function formatDates() {
      history.forEach(record => {
        record.date = new Date(record.date).toLocaleDateString('es-ES');
      });
    }
  
    fetchHistory();
  </script>
  
  
<style>
  table {
    width: 100%;
    border-collapse: collapse;
  }

  th, td {
    padding: 8px;
    text-align: left;
  }

  th {
    background-color: #f2f2f2;
  }

  tr:nth-child(even) {
    background-color: #f2f2f2;
  }

  .back-button {
    margin-top: 20px;
    cursor: pointer;
    padding: 10px 20px;
    background-color: #2b829c;
    color: rgb(219, 213, 213);
    border: none;
    border-radius: 5px;
    transition: background-color 0.3s ease;
  }

  .back-button:hover {
    background-color: #1d7a96;
  }
  </style>
  
  <h2>Historial de Operaciones</h2>
  
  {#if history.length > 0}
    <table>
      <thead>
        <tr>
          <th>Operación</th>
          <th>Resultado</th>
          <th>Fecha</th>
        </tr>
      </thead>
      <tbody>
        {#each history as record}
          <tr>
            <td>{record.operation}</td>
            <td>{record.result}</td>
            <td>{record.date}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  {:else}
    <p>No hay operaciones en el historial.</p>
  {/if}
  <button class="back-button" on:click={() => navigate("/")}>Volver a la Calculadora</button>
  