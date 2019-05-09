<template>

    <div id="crea">

        <h3  v-if="tipo=='crear'"> Crear Receta</h3>
        <h3  v-if="tipo=='actualizar'"> Crear Receta</h3>
        <b-input
            id="nom"
            type="text"
            placeholder="Nombre Receta"
            v-model="nombre"
            required            
        ></b-input>
        <Ingrediente v-for="(ingrediente,index) in ingredientes" 
        v-on:childToParent="actualizarComponente" :num="index" 
        v-bind:key="index+ingrediente.nombre+ingrediente.cantidad+ingrediente.unidad" 
        :nom="ingrediente.nombre" :can="ingrediente.cantidad" :uni="ingrediente.unidad"
        v-on:eliminar="eliminarIngrediente" />

        <b-button variant="outline-primary" v-on:click="agregarIngrediente" >Agregar ingrediente</b-button>
        <b-textarea
            type="text"
            placeholder="Preparacion"
            v-model="preparacion"
            required
        ></b-textarea>
        <b-button v-if="tipo=='crear'" variant="outline-primary"  v-on:click="crear">Crear receta</b-button>
        <b-button v-if="tipo=='actualizar'" variant="outline-primary"  v-on:click="actualizar">Actualizar receta</b-button>
        <b-button v-if="tipo=='actualizar'" variant="outline-primary"  
        v-on:click="$emit('cancelar')">cancelar</b-button>

    </div>
</template>

<script>



import Vue from 'vue'
import Ingrediente from './Ingrediente'
import VueResource from 'vue-resource'
Vue.use(VueResource);


export default {
    name: 'Crear', 
    components: {
        Ingrediente
    },

    data () {
        return{
            nombreInicial:'',
            preparacion:'',
            nombre:'',
            ingredientes:[{nombre:'',
                           cantidad:0,
                           unidad:'null'}]
        }
    },

    created: function(){
        this.preparacion=this.prep,
        this.nombreInicial=this.nom,
        this.nombre=this.nom,
        this.ingredientes=this.ing
    },

    props:{
        tipo:String,
        prep:{
            type: String,
            default: ''
        },
        nom:{
            type: String,
            default: ''
        },
        ing:{
            type: Array,
            default: () => []
        }
    },

    methods:{
        agregarIngrediente: function(){
            this.ingredientes.push({nombre:'',
                                        cantidad:0,
                                        unidad:'null'})            
            console.log(this.ingredientes)
        },

        eliminarIngrediente: function(e){
            console.log(e)
            console.log(this.ingredientes.splice(e, 1))
            this.ingredientes=this.ingredientes
            
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
    
            if(this.nombre.length!=0 && this.preparacion.length!=0 && this.ingredientesLlenos()){
                this.$http.post('http://localhost:3000/receta',
                        {preparacion:this.preparacion,
                        nombre:this.nombre,
                        ingredientes:this.ingredientes
                }).then(response => {
                    
                    this.$emit('creado')
                    alert("Receta creada exitosamente!")
                    
                }, response => {
                    alert("Hubo un problema al conectar con la base de datos")
                });
            }
            else{
                alert("Llenar todos los campos")
            }
        },

        ingredientesLlenos:function(){
            for (let i = 0; i < this.ingredientes.length; i++) {
                 if(this.ingredientes[i].nombre.length==0
                  || this.ingredientes[i].unidad.length==0)return false
            }
               
            return true
        },

        actualizar: function(){
            console.log( this.preparacion)
            console.log( this.nombreInicial)
            console.log( this.nombre)
            console.log( this.ingredientes)
            console.log("ASFFFFFFFFFFFFFFFFFF")

            if(this.nombre.length!=0 && this.preparacion.length!=0 && this.ingredientesLlenos()){
                this.$http.post('http://localhost:3000/receta/'+this.nombreInicial,
                        {preparacion:this.preparacion,
                        nombre:this.nombre,
                        ingredientes:this.ingredientes
                }).then(response => {
                
                    this.$emit('actualizado',this.nombre)
                    alert("Receta actualizada exitosamente")
                    
                }, response => {
                    alert("Hubo un problema al conectar con la base de datos")
                });
            }
            else{
                alert("Llenar todos los campos")
            }
        }
    }    
}

</script>
<style>

#crea{
    display: inline-block;
    padding: 5%;
    background-color: #ffec61;
    margin-left: 25%;
    margin-right: 25%;
    width: 50%;
    border-radius: 15px;
}

#nom{
    margin-bottom: 5%;
}

textarea{
    margin-bottom: 5%;
    margin-top: 5%;
    border-radius: 10px;
}



</style>
