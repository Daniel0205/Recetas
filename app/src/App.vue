<template >
  <div id="app">

    <h1 align="center">RECETAS</h1>
    <div id="botones">
      <b-button variant="outline-primary" id="crear" v-on:click="cambiarEstado">Crear receta</b-button>
      <b-button variant="outline-primary" id="consultar" v-on:click="cambiarEstado">Consultar receta</b-button>
    </div>
    <div id="cre" v-if="estado=='crear'">
      <Crear v-on:creado="estado='inicio'" :tipo="'crear'"></Crear>
    </div >
    <Consultar v-if="estado=='consultar'"></Consultar>
  </div> 
</template>

<script>
import Vue from 'vue'
import VSwitch from 'v-switch-case' 
import Crear from './components/Crear'
import Consultar from './components/Consulta'


Vue.use(VSwitch)

export default {
  name: 'app',
  components: {
    Crear,
    Consultar
  },
  data () {
    return{
      estado:'inicio',
    }
  },

  methods: {
    cambiarEstado: function (event){
      
      switch(event.target.id){
        case 'crear':
          this.estado='crear'
          break;
        case 'consultar':
          this.estado='consultar'
          break;
        default:
          this.estado='inicio'
          break;
      }      
    }
  }
}
</script>

<style>

#botones{
  display: flex;
  justify-content: center;
  padding-top: 1%;
}

#consultar{
  margin-inline-start: 15%;
}
h1, .h1 {
  background: black;
  color: white;
  font-size: 2.5rem;
  height: 100px;

}

.btn{
  background-color: white
}

#cre{
  text-align:center;
}

</style>
