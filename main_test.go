package main

import (
	"strings"
	"testing"
)

func TestProcesFile(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8" ?>
<xml>
<content_categories>
<content_category><id>1</id><title>Музыка</title><description>Ау11111111дио</description><hru>music</hru></content_category>
<content_category><id>14</id><title>Фильмы</title><description>Специальная категория(формат) для ivi.ru</description><hru>movies</hru></content_category>
<content_category><id>15</id><title>Сериалы</title><description/><hru>series</hru></content_category>
<content_category><id>16</id><title>Программы</title><description/><hru>programs</hru></content_category>
<content_category><id>17</id><title>Мультфильмы</title><description/><hru>animation</hru></content_category>
<content_category><id>18</id><title>Не определено</title><description/><hru>ne_opredeleno</hru></content_category>
<content_category><id>19</id><title>Документальное кино</title><description>Спец категория для реестра. Карточки контента должны иметь категорию "Кино"</description><hru>dokumentalnoe_kino</hru></content_category>
<content_category><id>20</id><title>Для детей</title><description>Контент для детей</description><hru>dlya_detej</hru></content_category>
<content_category><id>22</id><title>Основной раздел</title><description>Основной раздел</description><hru>general_partition</hru></content_category>
</content_categories></xml>`

	err := processFile(strings.NewReader(data), "testdata")
	if err != nil {
		t.Error(err)
	}
}
