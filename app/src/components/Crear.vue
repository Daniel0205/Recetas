<template>

    <div>
        <b-input
            placeholder="Nombre Receta"
            v-model="nombre"
            
        ></b-input>
        <Ingrediente v-for="(ingrediente,index) in ingredientes" 
        v-on:childToParent="actualizarComponente" :num="index" v-bind:key="index" />

        <b-button variant="outline-primary" v-on:click="agregarIngrediente" >Agregar ingrediente</b-button>
        <b-button variant="outline-primary" v-on:click="eliminarIngrediente" >Eliminar ingrediente</b-button>
        <b-input
            placeholder="Preparacion"
            v-model="preparacion"
        ></b-input>
        <b-button variant="outline-primary"  v-on:click="crear">Crear receta</b-button>

    </div>
</template>

<script>



import Vue from 'vue'
import Ingrediente from './ingrediente'
import VueResource from 'vue-resource'
Vue.use(VueResource);


export default {
    name: 'Crear', 
    components: {
        Ingrediente
    },

    data () {
        return{
            estado:'inicio',
            preparacion:'',
            nombre:'',
            ingredientes:[{nombre:'',
                           cantidad:0,
                           unidad:''}]
        }
    },



    methods:{
        agregarIngrediente: function(){
            this.ingredientes.push({nombre:'',
                                        cantidad:0,
                                        unidad:''})            
            console.log(this.ingredientes)
        },

        eliminarIngrediente: function(){
            this.ingredientes.pop()            
            console.log(this.ingredientes)
        },

        actualizarComponente:function(e){
            if(e.nom!=null){
                this.ingredientes[e.ident].nombre=e.nom
            }
            else if(e.cant!=null){
                this.ingredientes[e.ident].cantidad=e.cant
            }
            else if(e.unid!=null){
                this.ingredientes[e.ident].unidad=e.unid
            }

            console.log(this.ingredientes)
            console.log(e)
        },

        

        crear: function(){
            this.$http.post('http://localhost:3000/receta',
                    {preparacion:this.preparacion,
                    nombre:this.nombre,
                    ingredientes:this.ingredientes
                }).then(response => {
                console.log(response.Body)
                $emit('creado', {creado:true})
                alert("Si, sirvio!!! :D")
                
            }, response => {
                console.log(response.Body)
                alert("No, sirvio!!! :(")
            });
        }
    }    
}

</script>
<style>

</style>
