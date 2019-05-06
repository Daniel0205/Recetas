<template >
  <div id="app"  v-switch="estado">

    <b-button variant="outline-primary" id="crear" v-on:click="cambiarEstado">Crear receta</b-button>
    <b-button variant="outline-primary" id="borrar" v-on:click="cambiarEstado">Borrar receta</b-button>
    <b-button variant="outline-primary" id="actualizar" v-on:click="cambiarEstado">Actualizar receta</b-button>
    <b-button variant="outline-primary" id="consultar" v-on:click="cambiarEstado">Consultar receta</b-button>
    <div v-case="crear">
      <Crear v-on:creado="hola"></Crear>
    </div>
  </div> 
</template>

<script>
import Vue from 'vue'
import VSwitch from 'v-switch-case' 
import Crear from './components/Crear'


Vue.use(VSwitch)

export default {
  name: 'app',
  components: {
    Crear
  },
  data () {
    return{
      estado:'inicio',
    }
  },

  methods: {
    cambiarEstado: function (event){
      console.log(event)
      
      switch(event.target.id){
        case 'crear':
          this.estado='crear'
          break;
        case 'borrar':
          this.estado='borrar'
          break;
        case 'actualizar':
          this.estado='actualizar'
          break;
        case 'consultar':
          this.estado='consultar'
          break;
        default:
          this.estado='inicio'
          break;
      }      
    }, 
    
    hola: function(event){
      console.log(event)
      estado='inicio'
    }
  }
}
</script>

<style>

</style>
