<template>
    <div>
        <b-row>
            <b-col sm="3">
                <b-input
                v-model="nombre"
                placeholder="Nombre Ingrediente"
                v-on:keyup="$emit('childToParent', {nom:nombre, ident:num})"
                ></b-input>
            </b-col>

            <b-col sm="3">
                <b-input
                type="number"
                step="0.1"
                v-model="cantidad"
                placeholder="cantidad"
                v-on:keyup="$emit('childToParent', {cant:parseFloat(cantidad), ident:num})"
                v-on:change="$emit('childToParent', {cant:parseFloat(cantidad), ident:num})"
                ></b-input>
            </b-col>
            .
            <b-col sm="3">
                <b-form-select
                :options="{ 'Ltr': 'Ltr', 'gr': 'gr', 'kilo': 'kilo' }"
                v-model="unidad"
                :value="null"
                v-on:change="$emit('childToParent', {unid:unidad, ident:num})"
                >
                    <option slot="first" :value="null">Elegir unidad de medida</option>
                </b-form-select>
            </b-col>
        </b-row>
    </div>
</template>

<script>

export default {
    name: 'Crear', 
    data () {
        return{
            estado:'inicio',
            unidad:'null',
            nombre:'',
            cantidad:''
        }
    },

    props:{
        num:Number,
    }
    
}
</script>
<style>

</style>