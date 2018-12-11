" General {{{
set nocompatible " Enables us Vim specific features
set noswapfile

" Show line, column number, and relative position within a file in the status line
set ruler

set history=1024
set autochdir
set whichwrap=b,s,<,>,[,]
set nobomb
set autowrite                   " Automatically save before :next, :make etc.
set noerrorbells                " No beeps
set nobackup                    " Don't create annoying backup files
set backspace=indent,eol,start whichwrap+=<,>,[,]
" 设置 alt 键不映射到菜单栏
set winaltkeys=no
set mouse=
" }}}

" Lang & Encoding {{{
set fileencodings=utf-8,gbk2312,gbk,gb18030,cp936
set encoding=utf-8
let $LANG = 'en_US.UTF-8'
"language messages zh_CN.UTF-8
" }}}

" GUI {{{
set background=dark
set t_Co=256
let g:rehash256 = 1
let g:molokai_original = 1
colorscheme molokai
set cursorline
set hlsearch
set number
" 分割出来的窗口位于当前窗口下边/右边
set splitbelow
set splitright
" 使用内置 tab 样式而不是 gui
set guioptions-=e
set nolist
" set listchars=tab:▶\ ,eol:¬,trail:·,extends:>,precedes:<
set guifont=courier_new:h18    
" }}}

" Format {{{
set autoindent
set smartindent
set tabstop=4
set expandtab
set softtabstop=4
" set foldmethod=indent
syntax on
" }}}

" Keymap {{{
let mapleader=","

" 移动分割窗口
nmap <C-j> <C-W>j
nmap <C-k> <C-W>k
nmap <C-h> <C-W>h
nmap <C-l> <C-W>l

noremap <LEFT> <NOP>
noremap <RIGHT> <NOP>
noremap <Up> <NOP>
noremap <Down> <NOP>
noremap j gj
noremap k gk

" 正常模式下 alt+j,k,h,l 调整分割窗口大小
" nnoremap <M-j> :resize +5<cr>
" nnoremap <M-k> :resize -5<cr>
" nnoremap <M-h> :vertical resize -5<cr>
" nnoremap <M-l> :vertical resize +5<cr>

" 插入模式移动光标 alt + 方向键
inoremap <M-j> <Down>
inoremap <M-k> <Up>
inoremap <M-h> <left>
inoremap <M-l> <Right>

nnoremap <leader>w :set wrap!<cr>

imap <C-v> "+gP
vmap <C-c> "+y
vnoremap <BS> d
vnoremap <C-C> "+y
vnoremap <C-Insert> "+y
imap <C-V>		"+gP
map <S-Insert>		"+gP
cmap <C-V>		<C-R>+
cmap <S-Insert>		<C-R>+

" 复制当前文件/路径到剪贴板
nmap ,fn :let @*=substitute(expand("%"), "/", "\\", "g")<CR>
nmap ,fp :let @*=substitute(expand("%:p"), "/", "\\", "g")<CR>

" 设置切换Buffer快捷键"
nnoremap <C-left> :bn<CR>
nnoremap <C-right> :bp<CR>

" Shortcuts remap
nmap  <F2>          :tabnew<CR>
nmap  <F4>          :set hlsearch<CR>
map <F5> :NERDTreeToggle<CR>
nmap  <F7>          :GundoToggle<CR>
nmap  <F8>          :TagbarToggle<CR>
nmap  <F9>          :tabclose<CR>
"nmap  <F10>         :Tabularize /,\zs<CR>
"vmap  <F10>         :Tabularize /,\zs<CR>
"nmap  <F12>         :Tabularize / \zs<CR>
"vmap  <F12>         :Tabularize / \zs<CR>
map  <F10>      :tabprevious<CR>
map  <F11>        :tabnext<CR>

" }}}

"插件管理
call plug#begin('~/.vim/plugged')
" 可以在此次安装插件
Plug 'vim-airline/vim-airline'
Plug 'fatih/molokai'
Plug 'scrooloose/nerdtree'
Plug 'vim-airline/vim-airline-themes'
Plug 'Valloric/YouCompleteMe', {'tag': 'e018777'}
Plug 'fatih/vim-go'
" 代码提示
Plug 'Blackrush/vim-gocode'
" 边栏
Plug 'majutsushi/tagbar'
Plug 'ervandew/supertab'
Plug 'SirVer/ultisnips'
Plug 'dgryski/vim-godef'
call plug#end()
filetype plugin indent on

" YouCompleteMe settings
let g:ycm_collect_identifiers_from_comments_and_strings = 0
" 输入第1个字符开始补全 @2018-07-19 改为1，否则移动光标会弹出一堆无意义的提示
let g:ycm_min_num_of_chars_for_completion=1
" 禁止缓存匹配项,每次都重新生成匹配项
let g:ycm_cache_omnifunc=0
" 开启语义补全
let g:ycm_seed_identifiers_with_syntax=1
" 在注释输入中也能补全
let g:ycm_complete_in_comments = 1
" 在字符串输入中也能补全
let g:ycm_complete_in_strings = 1
let g:ycm_autoclose_preview_window_after_completion = 1
let g:ycm_autoclose_preview_window_after_insertion = 1
" let g:ycm_filetype_whitelist = { 'python': 1 }
let g:ycm_python_binary_path = 'python3'
map <C-G>  :YcmCompleter GoToDefinitionElseDeclaration<CR>
"  fix .cpp error: ValueError: Still no compile flags, no completions yet.
let g:ycm_global_ycm_extra_conf = '~/.vim/plugged/YouCompleteMe/third_party/ycmd/cpp/ycm/.ycm_extra_conf.py'
" Vim's autocomplete is excruciatingly slow
" http://stackoverflow.com/questions/2169645/vims-autocomplete-is-excruciatingly-slow
set complete-=i


let g:godef_same_file_in_same_window=1
let g:godef_split=2
let g:UltiSnipsExpandTrigger="<c-m>"
" vim-go
let g:go_fmt_command = "goimports"
let g:go_autodetect_gopath = 1
let g:go_list_type = "quickfix"

let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_generate_tags = 1

augroup go
  autocmd!

  " Show by default 4 spaces for a tab
  autocmd BufNewFile,BufRead *.go setlocal noexpandtab tabstop=4 shiftwidth=4

  " :GoBuild and :GoTestCompile
  autocmd FileType go nmap <leader>b :<C-u>call <SID>build_go_files()<CR>

  " :GoTest
  autocmd FileType go nmap <leader>t  <Plug>(go-test)

  " :GoRun
  autocmd FileType go nmap <leader>r  <Plug>(go-run)

  " :GoDoc
  autocmd FileType go nmap <Leader>d <Plug>(go-doc)

  " :GoCoverageToggle
  autocmd FileType go nmap <Leader>c <Plug>(go-coverage-toggle)

  " :GoInfo
  autocmd FileType go nmap <Leader>i <Plug>(go-info)

  " :GoMetaLinter
  autocmd FileType go nmap <Leader>l <Plug>(go-metalinter)

  " :GoDef but opens in a vertical split
  autocmd FileType go nmap <Leader>v <Plug>(go-def-vertical)
  " :GoDef but opens in a horizontal split
  autocmd FileType go nmap <Leader>s <Plug>(go-def-split)

  " :GoAlternate  commands :A, :AV, :AS and :AT
  autocmd Filetype go command! -bang A call go#alternate#Switch(<bang>0, 'edit')
  autocmd Filetype go command! -bang AV call go#alternate#Switch(<bang>0, 'vsplit')
  autocmd Filetype go command! -bang AS call go#alternate#Switch(<bang>0, 'split')
  autocmd Filetype go command! -bang AT call go#alternate#Switch(<bang>0, 'tabe')
augroup END


" build_go_files is a custom function that builds or compiles the test file.
" It calls :GoBuild if its a Go file, or :GoTestCompile if it's a test file
function! s:build_go_files()
  let l:file = expand('%')
  if l:file =~# '^\f\+_test\.go$'
    call go#test#Test(0, 1)
  elseif l:file =~# '^\f\+\.go$'
    call go#cmd#Build(0)
  endif
endfunction
