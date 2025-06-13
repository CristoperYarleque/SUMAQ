export function getEmbedUrl(url) {
  // Match URLs de YouTube estándar, cortos (shorts) y youtu.be
  const youtubeMatch = url.match(
    /(?:youtube\.com\/(?:watch\?v=|shorts\/)|youtu\.be\/)([^\s&?\/]+)/
  );
  if (youtubeMatch) {
    return `https://www.youtube.com/embed/${youtubeMatch[1]}`;
  }
  return url;
}