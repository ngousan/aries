<template>
  <div class="editEditor-box">
    <div id="editEditor"></div>
  </div>
</template>

<script>
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import { uploadImgToAttachment } from '@/api/aries/picture'

export default {
  name: 'editEditor',
  data () {
    return {
      contentEditor: null,
      headers: {},
      uploadUrl: ''
    }
  },
  props: {
    content: {
      type: String,
      default: ''
    }
  },
  mounted () {
    this.initEditor()
  },
  methods: {
    // 初始化 vditor
    initEditor () {
      this.contentEditor = new Vditor('editEditor', {
        height: 380,
        toolbarConfig: {
          pin: true
        },
        cache: {
          enable: false
        },
        hint: {
          emoji: {
            pray: '🙏',
            broken_heart: '💔',
            ok_hand: '👌',
            smile: '😄',
            laughing: '😆',
            smirk: '😏',
            heart_eyes: '😍',
            grin: '😁',
            stuck_out_tongue: '😛',
            expressionless: '😑',
            unamused: '😒',
            sob: '😭',
            joy: '😂',
            tired_face: '😫',
            blush: '😊',
            cry: '😢',
            fearful: '😨'
          }
        },
        counter: 100000,
        upload: {
          accept: '.jpg,.png,.gif,.jpeg',
          max: 2 * 1024 * 1024,
          filename: name =>
            name
              .replace(/[^(a-zA-Z0-9\u4e00-\u9fa5.)]/g, '')
              .replace(/[?\\/:|<>*[\]()$%{}@~]/g, '')
              .replace('/\\s/g', ''),
          // 自定义上传
          handler: files => {
            const formData = new FormData()
            for (const file of files) {
              formData.append('file[]', file)
            }
            uploadImgToAttachment(formData)
              .then(res => {
                const imgUrl = res.data.imgUrl
                this.insertContent(`![${imgUrl}](${imgUrl})`)
              })
              .catch(() => {
              })
          }
        },
        after: () => {
          this.setContent(this.md_content)
        }
      })
    },
    // 获取编辑器文本
    getContent () {
      return this.contentEditor.getValue()
    },
    // 获取 HTML
    getHTML () {
      // 获取 markdown 的 HTML 内容
      return this.contentEditor.getHTML()
    },
    // 设置编辑器文本
    setContent (val) {
      this.contentEditor.setValue(val)
    },
    // 插入文本
    insertMDContent (val) {
      this.contentEditor.insertValue(val)
    }
  }
}
</script>

<style lang="scss">
</style>
