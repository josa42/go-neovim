if exists('g:loaded_case')
    finish
endif
let g:loaded_case = 1
let s:plugin_root = fnamemodify(resolve(expand('<sfile>:p')), ':h:h')

function! s:StartPlugin(host) abort
  return jobstart(s:plugin_root.'/bin/case', {'rpc': v:true})
endfunction

call remote#host#Register('case', 'x', function('s:StartPlugin'))

call remote#host#RegisterPlugin('case', '0', [
\ {'type': 'function', 'name': 'Handler_b33d0c0da04fae2980d7f160c26bd50b', 'sync': 1, 'opts': {}},
\ {'type': 'function', 'name': 'OperatorFunc_b33d0c0da04fae2980d7f160c26bd50b', 'sync': 1, 'opts': {}},
\ ])
