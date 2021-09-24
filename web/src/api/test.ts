import Http from '../utils/http'
export default defineComponent({
  setup() {
    const store = useStore()
    const getData = () => {
      Http.get('url').then((res: Object) => {
        console.log(res)
      })
    }
    return {
      store,
      getData
    }
  }
})