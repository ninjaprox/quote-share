import QuoteShare from './src/QuoteShare.js';
import './style.css';

const popup = new QuoteShare({
  imageSelector: 'img.cover-image',
});

popup.init()
