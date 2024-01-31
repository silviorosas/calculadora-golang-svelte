
<script>
// @ts-nocheck
  import { Router, Route } from "svelte-navigator";
  import Calculadora from "./Calculadora.svelte";
  import Historial from "./ListarOperaciones.svelte";
  
</script>

<Router>  
  <Route path="/" component={Calculadora} />
  <Route path="/historial" component={Historial} />
</Router>
